package crawl

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"github.com/exepirit/yggmap/pkg/collection"
)

type NetworkVisitor interface {
	VisitNode(node network.Node) bool
	VisitLink(from, to network.PublicKey) bool
}

func WalkNetwork(ctx context.Context, client *adminapi.Client, visitor NetworkVisitor) error {
	visitedNodes := collection.NewSet[string]()
	visitQueue := collection.NewQueue[network.PublicKey]()

	getSelfResponse, err := client.GetSelf()
	if err != nil {
		return fmt.Errorf("failed to get current node: %w", err)
	}
	visitQueue.Put(network.MustParseKey(getSelfResponse.PublicKey))

	currentNodeKey, haveNode := visitQueue.Pop()
	for haveNode {
		currentNode, neighbors, err := retriveNodeInfo(ctx, client, currentNodeKey)
		if err != nil {
			return fmt.Errorf("cannot crawl node %q", currentNodeKey.String())
		}
		visitedNodes.Put(currentNode.PublicKey.String())

		mustContinue := visitor.VisitNode(*currentNode)
		if !mustContinue {
			continue
		}

		for _, neighbor := range neighbors {
			mustContinue = visitor.VisitLink(currentNode.PublicKey, neighbor)
			if !mustContinue {
				continue
			}

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
