package network

type Network struct {
	Nodes []Node
	Links []NodesLink
}

func (net *Network) AddNode(node Node, neighbors []PublicKey) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return
		}
	}
	net.Nodes = append(net.Nodes, node)

	for _, n := range neighbors {
		net.Links = append(net.Links, NodesLink{
			From: node.PublicKey,
			To:   n,
		})
	}
}

func (net Network) GetNode(key PublicKey) (Node, bool) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == key.String() {
			return n, true
		}
	}
	return Node{}, false
}
