package netstat

import (
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
)

func crawlCurrent(client *adminapi.Client) (yggdrasil.Node, error) {
	getSelfResponse, err := client.GetSelf()
	if err != nil {
		return yggdrasil.Node{}, err
	}

	publicKey, err := yggdrasil.ParseKey(getSelfResponse.PublicKey)
	if err != nil {
		return yggdrasil.Node{}, fmt.Errorf("invalid node key %q", getSelfResponse.PublicKey)
	}

	info, err := crawlAdditionalInfo(client, getSelfResponse.PublicKey)
	if err != nil {
		return yggdrasil.Node{}, fmt.Errorf("failed to additional node info: %w", err)
	}

	return yggdrasil.Node{
		PublicKey:      publicKey,
		AdditionalInfo: info,
	}, nil
}

func crawlNode(client *adminapi.Client, key yggdrasil.PublicKey) (yggdrasil.Node, error) {
	info, err := crawlAdditionalInfo(client, key.String())
	if err != nil {
		return yggdrasil.Node{}, fmt.Errorf("failed to additional node info: %w", err)
	}

	return yggdrasil.Node{
		PublicKey:      key,
		AdditionalInfo: info,
	}, nil
}

func crawlAdditionalInfo(client *adminapi.Client, key string) (map[string]any, error) {
	getInfoResponse, err := client.GetNodeInfo(key)
	if err != nil {
		return nil, err
	}
	info, _ := getInfoResponse[key]
	return info, nil
}

func crawlCurrentNodeNeighbors(client *adminapi.Client) ([]yggdrasil.PublicKey, error) {
	getPeersResponse, err := client.GetPeers()
	if err != nil {
		return nil, err
	}

	keys := make([]yggdrasil.PublicKey, 0, len(getPeersResponse.Peers))
	for _, peer := range getPeersResponse.Peers {
		key, err := yggdrasil.ParseKey(peer.PublicKey)
		if err != nil {
			return nil, fmt.Errorf("invalid node key %q", peer.PublicKey)
		}
		keys = append(keys, key)
	}
	return keys, nil
}

func crawlNodeNeighbors(client *adminapi.Client, key yggdrasil.PublicKey) ([]yggdrasil.PublicKey, error) {
	getPeersResponse, err := client.RemoteGetPeers(key.String())
	if err != nil {
		return nil, err
	}
	peers, ok := getPeersResponse[key.String()]
	if !ok {
		return []yggdrasil.PublicKey{}, nil
	}

	keys := make([]yggdrasil.PublicKey, 0, len(peers.Keys))
	for _, peerKey := range peers.Keys {
		key, err := yggdrasil.ParseKey(peerKey)
		if err != nil {
			return nil, fmt.Errorf("invalid node key %q", peerKey)
		}
		keys = append(keys, key)
	}
	return keys, nil
}
