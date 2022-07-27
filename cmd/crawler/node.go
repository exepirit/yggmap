package main

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

func (crawler NodeCrawler) GetNode(key string) (*network.Node, error) {
	info := &network.Node{
		LastSeen: time.Now(),
		IsActive: true,
	}

	selfInfo, err := crawler.Client.RemoteGetSelf(key)
	if err != nil {
		return info, fmt.Errorf("get basic node info: %w", err)
	}

	addr := network.MustParseKey(key).IPv6Address()
	info.PublicKey = network.MustParseKey(selfInfo[addr].PublicKey)
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

func (crawler NodeCrawler) GetRoot() (*network.Node, error) {
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
