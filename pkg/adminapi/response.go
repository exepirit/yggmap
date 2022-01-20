package adminapi

type rawResponse struct {
	Response map[string]interface{} `json:"response"`
	Status   string                 `json:"status"`
}

type NodeInfo struct {
	PublicKey   string `mapstructure:"key"`
	Port        uint8  `mapstructure:"port"`
	Coordinates []int  `mapstructure:"coords"`
	RemoteURI   string `mapstructure:"remote"`
}

type GetDHTResponse struct {
	DHT map[string]NodeInfo `mapstructure:"dht"`
}

type GetSelfResponse struct {
	PublicKey    string `mapstructure:"key"`
	BuildName    string `mapstructure:"build_name"`
	BuildVersion string `mapstructure:"build_version"`
	Coordinates  []int  `mapstructure:"coords"`
	Subnet       string `mapstructure:"subnet"`
}

type GetPeersResponse struct {
	Peers map[string]NodeInfo `json:"peers"`
}

type GetNodeInfoResponse map[string]map[string]interface{}

type RemoteGetPeersResponse map[string]struct {
	Keys []string `mapstructure:"keys"`
}

type RemoteGetSelfResponse map[string]struct {
	Coordinates string `mapstructure:"coords"`
	PublicKey   string `mapstructure:"key"`
}

type RemoteGetDHTResponse map[string]struct {
	Keys []string `mapstructure:"keys"`
}
