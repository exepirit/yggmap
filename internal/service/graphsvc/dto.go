package graphsvc

type GraphDto struct {
	Nodes []NodeDto         `json:"nodes"`
	Edges map[string]string `json:"edges"`
}

type NodeDto struct {
	ID    string  `json:"id"`
	Label string  `json:"label"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}
