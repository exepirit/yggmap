package v1

type NetworkDto struct {
	Nodes []NodeDto `json:"nodes"`
	Edges []EdgeDto `json:"edges"`
}

type EdgeDto struct {
	From string `json:"from"`
	To   string `json:"to"`
}
