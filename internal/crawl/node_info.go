package crawl

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/adminapi"
)

type NodeCrawler struct {
	Client *adminapi.Client
}

func (crawler NodeCrawler) GetNode(key network.PublicKey) (*network.Node, error) {
	info := &network.Node{
		LastSeen: time.Now(),
		IsActive: true,
	}

	getSelfResponse, err := crawler.Client.RemoteGetSelf(key.String())
	if err != nil {
		return nil, fmt.Errorf("get basic node info: %w", err)
	}

	selfAddress := key.IPv6Address()
	selfInfo, ok := getSelfResponse[selfAddress]
	if !ok {
		return nil, fmt.Errorf("response doesn't contain node itself info")
	}

	info.PublicKey, err = network.ParseKey(selfInfo.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("invalid node public key: %w", err)
	}

	info.Coordinates = parseCoordinatesFromStr(selfInfo.Coordinates)
	info.AdditionalInfo = make(map[string]interface{})

	detailInfo, err := crawler.Client.GetNodeInfo(key.String())
	if err != nil {
		return info, fmt.Errorf("get detail node info: %w", err)
	}
	for k, v := range detailInfo[selfAddress] {
		info.AdditionalInfo[k] = v
	}

	return info, nil
}

func (crawler NodeCrawler) GetPeersKeys(targetKey network.PublicKey) ([]network.PublicKey, error) {
	peers, err := crawler.Client.RemoteGetPeers(targetKey.String())
	if err != nil {
		return nil, err
	}

	peersKeys := peers[targetKey.IPv6Address()].Keys
	keys := make([]network.PublicKey, len(peersKeys))
	for i, k := range peersKeys {
		keys[i] = network.MustParseKey(k)
	}
	return keys, nil
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
