package network

import (
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
)

type Network struct {
	Nodes []yggdrasil.Node
	Links []NodesLink
}

func (net *Network) AddNode(node yggdrasil.Node) error {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return fmt.Errorf("node already exists")
		}
	}
	net.Nodes = append(net.Nodes, node)
	return nil
}

func (net *Network) ConnectNodes(from, to yggdrasil.PublicKey) {
	net.Links = append(net.Links, NodesLink{
		From: from,
		To:   to,
	})
}

func (net *Network) GetNode(key yggdrasil.PublicKey) (yggdrasil.Node, bool) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == key.String() {
			return n, true
		}
	}
	return yggdrasil.Node{}, false
}
