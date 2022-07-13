package network

type Graph struct {
	Nodes []GraphNode
	Edges []Edge
}

func (graph Graph) GetNode(id string) *GraphNode {
	for i := range graph.Nodes {
		if graph.Nodes[i].GetID() == id {
			return &graph.Nodes[i]
		}
	}
	return nil
}

func (graph Graph) GetNeighbors(id string) []string {
	neighbors := make([]string, 0)
	for _, e := range graph.Edges {
		if e.From.String() == id {
			neighbors = append(neighbors, e.To.String())
		}
	}
	return neighbors
}

type GraphNode struct {
	Node *Node
	X, Y float64
}

func (node GraphNode) GetID() string {
	return node.Node.PublicKey.String()
}

func (node GraphNode) Size() float64 {
	size := float64(len(node.Node.Peers))
	if size < 1 {
		size = 1
	}
	return size
}
