package network

import "fmt"

type Network struct {
	Nodes []Node
	Links []NodesLink
}

func (net *Network) AddNode(node Node) error {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return fmt.Errorf("node already exists")
		}
	}
	net.Nodes = append(net.Nodes, node)
	return nil
}

func (net *Network) ConnectNodes(from, to PublicKey) {
	net.Links = append(net.Links, NodesLink{
		From: from,
		To:   to,
	})
}

func (net *Network) GetNode(key PublicKey) (Node, bool) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == key.String() {
			return n, true
		}
	}
	return Node{}, false
}
