package network

import (
	"github.com/exepirit/yggmap/internal/domain/node"
)

type Network struct {
	Nodes []*node.Node
	Edges []Edge
}

func (net *Network) AddConnection(from, to *node.Node) {
	net.AddNode(from)
	net.AddNode(to)

	net.Edges = append(net.Edges, Edge{
		From: from.PublicKey,
		To:   to.PublicKey,
	})
}

func (net *Network) AddNode(node *node.Node) {
	for _, n := range net.Nodes {
		if n.PublicKey.String() == node.PublicKey.String() {
			return
		}
	}
	net.Nodes = append(net.Nodes, node)
}
