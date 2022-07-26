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

func (crawler *NetworkCrawler) GetNetwork(ctx context.Context) (network.Network, error) {
	net := network.Network{}

	root, err := crawler.node().GetRoot()
	if err != nil {
		return net, fmt.Errorf("cannot crawl root node: %w", err)
	}

	err = crawler.crawlRecursive(ctx, &net, root)
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

		node, peers, err := crawler.crawlOneNode(key)
		if err != nil {
			log.Warn().
				Str("nodeKey", key).Err(err).
				Msg("Cannot collect node")
			continue
		}
		scrapedNodes[key] = node
		log.Info().Str("nodeKey", key).
			Msgf("Scraped node %s", node.PublicKey.IPv6Address())

		for _, peer := range peers {
			peerKey := peer.String()
			pending := scrapeQueue.contains(peerKey)
			_, crawled := scrapedNodes[peerKey]
			crawled = crawled || (peerKey == key)
			if !pending && !crawled {
				scrapeQueue.push(peerKey)
			} else {
				log.Debug().Msgf("Ignored node %s. It already queued or scanned.", peer)
			}
		}

		net.AddNode(*node, peers)
		log.Debug().Msgf("%d nodes in crawl queue, %d scanned", scrapeQueue.length(), len(scrapedNodes))
	}

	// crawled network may have links to nodes which destroyed during
	// crawling it must be removed
	removeDanglingLinks(net)
	removeOrphanNodes(net)

	return nil
}

func (crawler *NetworkCrawler) crawlOneNode(targetKey string) (*network.Node, []network.PublicKey, error) {
	node, err := crawler.node().GetNode(targetKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get info: %w", err)
	}

	peers, err := crawler.node().GetPeersKeys(network.MustParseKey(targetKey))
	if err != nil {
		log.Warn().
			Str("nodeKey", targetKey).Err(err).
			Msg("Cannot get node peers. Part of network might not be scanned")
	}
	log.Debug().Msgf("Node %s have %d peers", targetKey, len(peers))

	// node may have > 1 connection with other node, but it must not be displayed on map
	peers = crawler.deduplicateKeys(peers)

	return node, peers, nil
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

func removeDanglingLinks(net *network.Network) {
	withoutDangling := make([]network.NodesLink, 0, len(net.Links))
	for _, link := range net.Links {
		_, srcExists := net.GetNode(link.From)
		_, dstExists := net.GetNode(link.To)
		if srcExists && dstExists {
			withoutDangling = append(withoutDangling, link)
		} else {
			log.Debug().Msg("Found dangling link. It removed from network")
		}
	}
	net.Links = withoutDangling
}

func removeOrphanNodes(net *network.Network) {
	haveNeighbors := func(node network.Node) bool {
		addr := node.PublicKey.String()
		for _, link := range net.Links {
			if link.From.String() == addr || link.To.String() == addr {
				return true
			}
		}
		return false
	}

	withoutOrphans := make([]network.Node, 0, len(net.Nodes))
	for _, node := range net.Nodes {
		if haveNeighbors(node) {
			withoutOrphans = append(withoutOrphans, node)
		} else {
			log.Debug().Msgf("Node %s is orphan. It excluded from network", node.PublicKey)
		}
	}
	net.Nodes = withoutOrphans
}
