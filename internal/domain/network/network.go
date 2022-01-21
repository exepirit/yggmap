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
	net.Nodes = append(net.Nodes, node)
}
