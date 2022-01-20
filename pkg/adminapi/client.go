package adminapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net"
)

type Client struct {
	socketType, address string
	conn                net.Conn
}

func (client *Client) Close() error {
	if client.conn != nil {
		return client.conn.Close()
	}
	return nil
}

func (client *Client) connection() (net.Conn, error) {
	if client.conn == nil {
		conn, err := net.Dial(client.socketType, client.address)
		if err != nil {
			return nil, err
		}
		client.conn = conn
	}
	return client.conn, nil
}

func (client *Client) doRequest(req interface{}) (rawResponse, error) {
	conn, err := client.connection()
	if err != nil {
		return rawResponse{}, fmt.Errorf("cannot connect to node: %w", err)
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	if err = encoder.Encode(req); err != nil {
		return rawResponse{}, fmt.Errorf("cannot send request: %w", err)
	}

	var resp rawResponse
	if err = decoder.Decode(&resp); err != nil {
		return rawResponse{}, fmt.Errorf("cannot read response: %w", err)
	}

	if resp.Status == "error" {
		reason, ok := resp.Response["error"]
		if ok {
			return resp, fmt.Errorf("%v", reason)
		}
		return resp, fmt.Errorf("internal server error")
	}

	return resp, nil
}

func (client *Client) GetSelf() (GetSelfResponse, error) {
	resp, err := client.doRequest(rawRequest("getSelf"))
	if err != nil {
		return GetSelfResponse{}, err
	}

	self := resp.Response["self"].(map[string]interface{})
	for _, nodeInfo := range self {
		nodeInfo := nodeInfo.(map[string]interface{})
		var result GetSelfResponse
		return result, mapstructure.Decode(nodeInfo, &result)
	}
	return GetSelfResponse{}, errors.New("response not contains node info (malformed response?)")
}

func (client *Client) GetPeers() (GetPeersResponse, error) {
	resp, err := client.doRequest(rawRequest("getPeers"))
	if err != nil {
		return GetPeersResponse{}, err
	}

	var result GetPeersResponse
	return result, mapstructure.Decode(resp.Response, &result)
}

func (client *Client) GetDHT() (GetDHTResponse, error) {
	resp, err := client.doRequest(rawRequest("getDHT"))
	if err != nil {
		return GetDHTResponse{}, err
	}

	var result GetDHTResponse
	return result, mapstructure.Decode(resp.Response, &result)
}

func (client *Client) GetNodeInfo(key string) (GetNodeInfoResponse, error) {
	resp, err := client.doRequest(requestWithKey("getNodeInfo", key))
	if err != nil {
		return GetNodeInfoResponse{}, err
	}

	var result GetNodeInfoResponse
	return result, mapstructure.Decode(resp.Response, &result)
}

func (client *Client) RemoteGetPeers(key string) (RemoteGetPeersResponse, error) {
	resp, err := client.doRequest(
		requestWithKey("debug_remoteGetPeers", key))
	if err != nil {
		return RemoteGetPeersResponse{}, err
	}

	var result RemoteGetPeersResponse
	return result, mapstructure.Decode(resp.Response, &result)
}

func (client *Client) RemoteGetDHT(key string) (RemoteGetDHTResponse, error) {
	resp, err := client.doRequest(
		requestWithKey("debug_remoteGetDHT", key))
	if err != nil {
		return RemoteGetDHTResponse{}, err
	}

	var result RemoteGetDHTResponse
	return result, mapstructure.Decode(resp.Response, &result)
}

func (client *Client) RemoteGetSelf(key string) (RemoteGetSelfResponse, error) {
	resp, err := client.doRequest(
		requestWithKey("debug_remoteGetSelf", key))
	if err != nil {
		return RemoteGetSelfResponse{}, err
	}

	var result RemoteGetSelfResponse
	return result, mapstructure.Decode(resp.Response, &result)
}
