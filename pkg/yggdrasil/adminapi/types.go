package adminapi

import (
	"github.com/yggdrasil-network/yggdrasil-go/src/admin"
)

// aliases to mainline admin api DTOs
type (
	Response         = admin.AdminSocketResponse
	GetSelfResponse  = admin.GetSelfResponse
	GetPeersResponse = admin.GetPeersResponse
)

type Request struct {
	Name      string            `json:"request"`
	Arguments map[string]string `json:"arguments"`
	KeepAlive bool              `json:"keepalive"`
}

type GetNodeInfoResponse map[string]map[string]interface{}

type RemoteGetPeersResponse map[string]struct {
	Keys []string `json:"keys"`
}

type RemoteGetSelfResponse map[string]struct {
	PublicKey string `json:"key"`
}
