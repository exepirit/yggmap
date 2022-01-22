package network

type Network struct {
	Nodes []*Node
}

func (net *Network) AddNode(node *Node) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return
		}
	}
}

func (net *Network) GetEdges() []Edge {
	result := make([]Edge, 0)
	for _, node := range net.Nodes {
		for _, peer := range node.Peers {
			result = append(result, Edge{
				From: node.PublicKey,
				To:   peer,
			})
		}
	}
	return result
}
