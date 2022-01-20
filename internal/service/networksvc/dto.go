package networksvc

type NetworkDto struct {
	Nodes []NodeDto `json:"nodes"`
	Edges []EdgeDto `json:"edges"`
}

type NodeDto struct {
	PublicKey      string                 `json:"publicKey"`
	AdditionalInfo map[string]interface{} `json:"additionalInfo"`
}

type EdgeDto struct {
	From string `json:"from"`
	To   string `json:"to"`
}
