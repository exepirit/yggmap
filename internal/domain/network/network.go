package network

type Network struct {
	Nodes []*Node
	Edges []*Edge
}

func (net *Network) AddNode(node *Node) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return
		}
	}

	net.Nodes = append(net.Nodes, node)
	for _, peer := range node.Peers {
		net.Edges = append(net.Edges, &Edge{
			From: node.PublicKey,
			To:   peer,
		})
	}
}
