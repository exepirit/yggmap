package crawl

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"github.com/exepirit/yggmap/pkg/collection"
)

// NetworkVisitor is an interface that allows users to visit nodes and links in a network.
type NetworkVisitor interface {
	// VisitNode visits a node in the network.
	// Returns true if the visitor should continue visiting the graph deep, false otherwise.
	VisitNode(node network.Node) bool

	// VisitLink visits a link between two nodes in the network.
	// Returns true if the visitor should continue visiting the graph deep, false otherwise.
	VisitLink(from, to network.PublicKey) bool
}

// WalkNetwork walks through a network and visits each node and link using a NetworkVisitor.
func WalkNetwork(ctx context.Context, client *adminapi.Client, visitor NetworkVisitor) error {
	visitedNodes := collection.NewSet[string]()
	visitQueue := collection.NewQueue[network.PublicKey]()

	// Get the current node and add it to the queue.
	getSelfResponse, err := client.GetSelf()
	if err != nil {
		return fmt.Errorf("failed to get current node: %w", err)
	}
	visitQueue.Put(network.MustParseKey(getSelfResponse.PublicKey))

	// Visit each node in the queue and its neighbors.
	currentNodeKey, haveNode := visitQueue.Pop()
	for haveNode {
		currentNode, neighbors, err := retriveNodeInfo(ctx, client, currentNodeKey)
		if err != nil {
			return fmt.Errorf("cannot crawl node %q", currentNodeKey.String())
		}
		visitedNodes.Put(currentNode.PublicKey.String())

		// Visit the current node using the NetworkVisitor.
		mustContinue := visitor.VisitNode(*currentNode)
		if !mustContinue {
			continue
		}

		for _, neighbor := range neighbors {
			// Visit the link between the current node and its neighbor using the NetworkVisitor.
			mustContinue = visitor.VisitLink(currentNode.PublicKey, neighbor)
			if !mustContinue {
				continue
			}

			// If the neighbor has not been visited before, add it to the queue.
			if !visitedNodes.Contains(neighbor.String()) {
				visitQueue.Put(neighbor)
			}
		}
	}

	return nil
}

func retriveNodeInfo(
	_ context.Context,
	yggdrasil *adminapi.Client,
	key network.PublicKey,
) (node *network.Node, peers []network.PublicKey, err error) {
	nodeInfoCrawler := NodeCrawler{Client: yggdrasil}
	node, err = nodeInfoCrawler.GetNode(key)
	if err != nil {
		return
	}
	peers, err = nodeInfoCrawler.GetPeersKeys(key)
	return
}
