package netstat

import (
	"errors"
	"fmt"
	"github.com/exepirit/yggmap/pkg/collection"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"log/slog"
)

// ErrStopIteration is returned when the walker has been asked to stop iterating over nodes.
// This might be due to an error occurring during node visit.
var ErrStopIteration = errors.New("stop iteration")

// Walker is a struct used to crawl through the yggdrasil network starting from local nodes. It uses a [Visitor] interface
// to handle the visited nodes and links, an [adminapi.Client] for node information gathering.
type Walker struct {
	Visitor Visitor
	Client  *adminapi.Client

	queue collection.Queue[yggdrasil.PublicKey]
	found collection.Set[string]
}

// StartFromLocal initiates the crawling process from local nodes. It first visits the local node, then recursively
// visits all reachable nodes starting from the local one.
func (w *Walker) StartFromLocal() error {
	err := w.visitLocal()
	if err != nil {
		return err
	}

	return w.visitRecursive()
}

func (w *Walker) visitLocal() error {
	node, neighbors, err := w.collectLocalNode()
	if err != nil {
		return fmt.Errorf("failed to crawl local node info: %w", err)
	}

	if next := w.Visitor.VisitNode(node); !next {
		return ErrStopIteration
	}
	w.found.Put(node.PublicKey.String())

	for _, neighborKey := range neighbors {
		if next := w.Visitor.VisitLink(node.PublicKey, neighborKey); !next {
			return ErrStopIteration
		}
		w.found.Put(neighborKey.String())
		w.queue.Put(neighborKey)
	}
	return nil
}

func (w *Walker) visitRecursive() error {
	// This ensures we keep visiting nodes as long as they're reachable from the local one.
	for w.queue.Len() > 0 {
		slog.Debug("Next network crawling iteration starts",
			"queueLength", w.queue.Len(), "foundCount", w.found.Len())

		nodeKey, _ := w.queue.Pop()
		slog.Debug("Collecting node information", "key", nodeKey.String())

		node, neighbors, err := w.collectNode(nodeKey)
		if err != nil {
			slog.Warn("Cannot crawl node", "key", nodeKey.String(), "error", err)
			continue
		}

		if next := w.Visitor.VisitNode(node); !next {
			return ErrStopIteration
		}

		for _, neighborKey := range neighbors {
			if next := w.Visitor.VisitLink(node.PublicKey, neighborKey); !next {
				return ErrStopIteration
			}

			// If the neighbor has not been seen before, we mark it as found and add
			// it to the queue for further visiting
			keyStr := neighborKey.String()
			if !w.found.Contains(keyStr) {
				w.found.Put(keyStr)
				w.queue.Put(neighborKey)
			}
		}
	}
	return nil
}

func (w *Walker) collectNode(key yggdrasil.PublicKey) (node yggdrasil.Node, neighbors []yggdrasil.PublicKey, err error) {
	node, err = crawlNode(w.Client, key)
	if err != nil {
		return
	}
	neighbors, err = crawlNodeNeighbors(w.Client, key)
	return
}

func (w *Walker) collectLocalNode() (node yggdrasil.Node, neighbors []yggdrasil.PublicKey, err error) {
	node, err = crawlCurrent(w.Client)
	if err != nil {
		return
	}
	neighbors, err = crawlCurrentNodeNeighbors(w.Client)
	return
}
