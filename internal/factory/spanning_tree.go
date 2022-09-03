package factory

import (
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/graph"
)

func MakeSpanningTree(nodes []network.Node) network.Network {
	net := network.Network{
		Nodes: nodes,
		Links: make([]network.NodesLink, 0, len(nodes)),
	}

	if len(nodes) < 1 {
		return net
	}

	tree := graph.MakeSpanningTree(nodes[0], nodes[1:], func(node network.Node) []int {
		return node.Coordinates
	})
	graph.WalkDepth(tree, func(parent *network.Node, child network.Node) {
		if parent == nil {
			return
		}

		net.Links = append(net.Links, network.NodesLink{
			From: parent.PublicKey,
			To:   child.PublicKey,
		})
	})

	return net
}
