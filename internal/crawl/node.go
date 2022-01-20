package crawl

import (
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/node"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"strconv"
	"strings"
)

type NodeCrawler struct {
	Client *adminapi.Client
}

func (crawler NodeCrawler) GetNode(key string) (*node.Node, error) {
	info := &node.Node{}
	selfInfo, err := crawler.Client.RemoteGetSelf(key)
	if err != nil {
		return info, fmt.Errorf("get basic node info: %w", err)
	}
	addr := node.MustParseKey(key).IPv6Address()
	info.PublicKey = node.MustParseKey(selfInfo[addr].PublicKey)
	info.Coordinates = parseCoordinatesFromStr(selfInfo[addr].Coordinates)
	info.AdditionalInfo = make(map[string]interface{})

	detailInfo, err := crawler.Client.GetNodeInfo(key)
	if err != nil {
		return info, fmt.Errorf("get detail node info: %w", err)
	}
	for k, v := range detailInfo[addr] {
		info.AdditionalInfo[k] = v
	}

	return info, nil
}

func (crawler NodeCrawler) GetPeersKeys(targetKey string) ([]string, error) {
	peers, err := crawler.Client.RemoteGetPeers(targetKey)
	if err != nil {
		return nil, err
	}
	addr := node.MustParseKey(targetKey).IPv6Address()
	return peers[addr].Keys, nil
}

func (crawler NodeCrawler) GetRoot() (*node.Node, error) {
	selfInfo, err := crawler.Client.GetSelf()
	if err != nil {
		return nil, err
	}
	return crawler.GetNode(selfInfo.PublicKey)
}

func parseCoordinatesFromStr(s string) []int {
	s = s[1 : len(s)-1]
	split := strings.Split(s, " ")
	coords := make([]int, len(split))
	for i, s := range split {
		coords[i], _ = strconv.Atoi(s)
	}
	return coords
}
