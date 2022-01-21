package crawl

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
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
func (crawler *NetworkCrawler) crawlRecursive(_ context.Context, net *network.Network, root *network.Node) error {
	scrapedNodes := make(map[string]*network.Node)
	scrapeQueue := newQueue()
	scrapeQueue.push(root.PublicKey.String())

	for !scrapeQueue.isEmpty() && len(scrapedNodes) < 10 {
		key, _ := scrapeQueue.pop()
		if _, scraped := scrapedNodes[key]; scraped {
			continue
		}

		node, err := crawler.crawlOneNode(key)
		if err != nil {
			log.Warn().
				Str("nodeKey", key).Err(err).
				Msg("Cannot collect scrapedNode")
			continue
		}
		scrapedNodes[key] = node
		log.Info().Str("nodeKey", key).
			Msgf("Scraped scrapedNode %s", node.PublicKey.IPv6Address())

		for _, peer := range node.Peers {
			peerKey := peer.String()
			pending := scrapeQueue.contains(peerKey)
			_, crawled := scrapedNodes[peerKey]
			crawled = crawled || (peerKey == key)
			if !pending && !crawled {
				scrapeQueue.push(peerKey)
			} else {
				log.Debug().Msgf("Ignored scrapedNode %s. It already queued or scanned.", peer)
			}
		}

		net.AddNode(node)
		log.Debug().Msgf("%d nodes in crawl queue, %d scanned", scrapeQueue.length(), len(scrapedNodes))
	}

	return nil
}

func (crawler *NetworkCrawler) crawlOneNode(targetKey string) (*network.Node, error) {
	node, err := crawler.node().GetNode(targetKey)
	if err != nil {
		return nil, fmt.Errorf("get info: %w", err)
	}

	nodePeers, err := crawler.node().GetPeersKeys(network.MustParseKey(targetKey))
	if err != nil {
		log.Warn().
			Str("nodeKey", targetKey).Err(err).
			Msg("Cannot get node peers. Part of network might not be scanned")
	}
	log.Debug().Msgf("Node %s have %d peers", targetKey, len(nodePeers))

	// node may have > 1 connection with other node, but it must not be displayed on map
	nodePeers = crawler.deduplicateKeys(nodePeers)

	return node, nil
}

func (NetworkCrawler) deduplicateKeys(keys []network.PublicKey) []network.PublicKey {
	m := make(map[string]network.PublicKey)
	for _, key := range keys {
		keyStr := key.String()
		if _, exists := m[keyStr]; !exists {
			m[keyStr] = key
		}
	}

	keys = make([]network.PublicKey, len(m))
	i := 0
	for _, key := range m {
		keys[i] = key
		i++
	}
	return keys
}
