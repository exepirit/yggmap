package netstat

import (
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"time"
)

type NodeCrawler struct {
	Client *adminapi.Client
}

func (crawler NodeCrawler) GetNode(key yggdrasil.PublicKey) (*yggdrasil.Node, error) {
	info := &yggdrasil.Node{
		LastSeen: time.Now(),
		IsActive: true,
	}

	// FIXME(bug): any debug_remoteget* method not working on local node
	getSelfResponse, err := crawler.Client.RemoteGetSelf(key.String())
	if err != nil {
		return nil, fmt.Errorf("get basic node info: %w", err)
	}

	selfAddress := key.IPv6Address()
	selfInfo, ok := getSelfResponse[selfAddress]
	if !ok {
		return nil, fmt.Errorf("response doesn't contain node itself info")
	}

	info.PublicKey, err = yggdrasil.ParseKey(selfInfo.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("invalid node public key: %w", err)
	}

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

func (crawler NodeCrawler) GetPeersKeys(targetKey yggdrasil.PublicKey) ([]yggdrasil.PublicKey, error) {
	peers, err := crawler.Client.RemoteGetPeers(targetKey.String())
	if err != nil {
		return nil, err
	}

	peersKeys := peers[targetKey.IPv6Address()].Keys
	keys := make([]yggdrasil.PublicKey, len(peersKeys))
	for i, k := range peersKeys {
		keys[i] = yggdrasil.MustParseKey(k)
	}
	return keys, nil
}
