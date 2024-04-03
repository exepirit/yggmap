package adminapi

import (
	"encoding/json"
	"fmt"
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

func (client *Client) connect() (net.Conn, error) {
	if client.conn == nil {
		conn, err := net.Dial(client.socketType, client.address)
		if err != nil {
			return nil, err
		}
		client.conn = conn
	}
	return client.conn, nil
}

func (client *Client) doRequest(request Request, response any) error {
	conn, err := client.connect()
	if err != nil {
		return fmt.Errorf("cannot connect to node: %w", err)
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	if err = encoder.Encode(request); err != nil {
		return fmt.Errorf("cannot send request: %w", err)
	}
	var rawResp Response
	if err = decoder.Decode(&rawResp); err != nil {
		return fmt.Errorf("cannot read response: %w", err)
	}

	if rawResp.Status == "error" {
		return fmt.Errorf("server error: %s", rawResp.Error)
	}

	if err = json.Unmarshal(rawResp.Response, response); err != nil {
		return fmt.Errorf("invalid server response: %w", err)
	}

	return nil
}

func (client *Client) GetSelf() (GetSelfResponse, error) {
	request := Request{Name: "getSelf", KeepAlive: true}
	var response GetSelfResponse
	return response, client.doRequest(request, &response)
}

func (client *Client) GetPeers() (GetPeersResponse, error) {
	request := Request{Name: "getPeers", KeepAlive: true}
	var response GetPeersResponse
	return response, client.doRequest(request, &response)
}

func (client *Client) GetNodeInfo(key string) (GetNodeInfoResponse, error) {
	request := Request{
		Name: "getNodeInfo",
		Arguments: map[string]string{
			"key": key,
		},
		KeepAlive: true,
	}
	var response GetNodeInfoResponse
	return response, client.doRequest(request, &response)
}

func (client *Client) RemoteGetPeers(key string) (RemoteGetPeersResponse, error) {
	request := Request{
		Name: "debug_remoteGetPeers",
		Arguments: map[string]string{
			"key": key,
		},
		KeepAlive: true,
	}
	var response RemoteGetPeersResponse
	return response, client.doRequest(request, &response)
}

func (client *Client) RemoteGetSelf(key string) (RemoteGetSelfResponse, error) {
	request := Request{
		Name: "debug_remoteGetSelf",
		Arguments: map[string]string{
			"key": key,
		},
		KeepAlive: true,
	}
	var response RemoteGetSelfResponse
	return response, client.doRequest(request, &response)
}
