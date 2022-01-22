package network

type Graph struct {
	Nodes []GraphNode
	Edges map[string]string
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
