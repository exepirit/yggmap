package crawl

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/internal/domain/node"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"github.com/rs/zerolog/log"
)

type NetworkCrawler struct {
	Client *adminapi.Client

	oneNode *NodeCrawler
}

func (crawler *NetworkCrawler) GetNetwork(ctx context.Context) (*network.Network, error) {
	net := &network.Network{}

	root, err := crawler.node().GetRoot()
	if err != nil {
		return nil, fmt.Errorf("cannot crawl root node: %w", err)
	}

	err = crawler.crawlRecursive(ctx, net, root)
	return net, err
}

func (crawler *NetworkCrawler) node() *NodeCrawler {
	if crawler.oneNode == nil {
		crawler.oneNode = &NodeCrawler{Client: crawler.Client}
	}
	return crawler.oneNode
}

// TODO: use context to cancel crawling process
func (crawler *NetworkCrawler) crawlRecursive(_ context.Context, net *network.Network, root *node.Node) error {
	scrapedNodes := make(map[string]nodeInfo)
	scrapeQueue := newQueue()
	scrapeQueue.push(root.PublicKey.String())

	for !scrapeQueue.isEmpty() && len(scrapedNodes) < 10 {
		key, _ := scrapeQueue.pop()
		if _, scraped := scrapedNodes[key]; scraped {
			continue
		}

		nodeFullInfo, err := crawler.crawlOneNode(key)
		if err != nil {
			log.Warn().
				Str("nodeKey", key).Err(err).
				Msg("Cannot collect scrapedNode")
			continue
		}
		scrapedNodes[key] = nodeFullInfo
		log.Info().Str("nodeKey", key).
			Msgf("Scraped scrapedNode %s", nodeFullInfo.PublicKey.IPv6Address())

		for _, peer := range nodeFullInfo.peers {
			pending := scrapeQueue.contains(peer)
			_, crawled := scrapedNodes[peer]
			crawled = crawled || (peer == key)
			if !pending && !crawled {
				scrapeQueue.push(peer)
			} else {
				log.Debug().Msgf("Ignored scrapedNode %s. It already queued or scanned.", peer)
			}
		}

		log.Debug().Msgf("%d nodes in crawl queue, %d scanned", scrapeQueue.length(), len(scrapedNodes))
	}

	for _, scrapedNode := range scrapedNodes {
		net.AddNode(scrapedNode.Node)
		for _, peerKey := range scrapedNode.peers {
			peer, ok := scrapedNodes[peerKey]
			if !ok {
				log.Warn().Str("nodeKey", peerKey).
					Msg("Couldn't find node in scraped nodes map. Part of network might not be displayed")
				continue
			}
			net.AddConnection(scrapedNode.Node, peer.Node)
		}
	}
	return nil
}

func (crawler *NetworkCrawler) crawlOneNode(targetKey string) (nodeInfo, error) {
	nodeBase, err := crawler.node().GetNode(targetKey)
	if err != nil {
		return nodeInfo{}, fmt.Errorf("get info: %w", err)
	}

	nodePeers, err := crawler.node().GetPeersKeys(targetKey)
	if err != nil {
		log.Warn().
			Str("nodeKey", targetKey).Err(err).
			Msg("Cannot get node peers. Part of network might not be scanned")
	}
	log.Debug().Msgf("Node %s have %d peers", targetKey, len(nodePeers))

	// node may have > 1 connection with other node, but it must not be displayed on map
	nodePeers = crawler.deduplicateKeys(nodePeers)

	return nodeInfo{
		Node:  nodeBase,
		peers: nodePeers,
	}, nil
}

func (NetworkCrawler) deduplicateKeys(keys []string) []string {
	m := make(map[string]bool)
	for _, key := range keys {
		if _, exists := m[key]; !exists {
			m[key] = true
		}
	}

	keys = make([]string, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	return keys
}

type nodeInfo struct {
	*node.Node
	peers []string
}
