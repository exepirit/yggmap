package factory

import "github.com/exepirit/yggmap/internal/domain/network"

func MakeGraphForNetwork(net *network.Network) *network.Graph {
	graph := &network.Graph{
		Nodes: make([]network.GraphNode, len(net.Nodes)),
		Edges: net.GetEdges(),
	}

	for i, node := range net.Nodes {
		graph.Nodes[i] = network.GraphNode{
			Node: node,
		}
	}

	return graph
}
