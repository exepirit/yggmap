package factory

import "github.com/exepirit/yggmap/internal/domain/network"

func MakeGraphForNetwork(net *network.Network) *network.Graph {
	graph := &network.Graph{
		Nodes: make([]network.GraphNode, len(net.Nodes)),
		Edges: make(map[string]string),
	}

	for i, node := range net.Nodes {
		graph.Nodes[i] = network.GraphNode{
			Node: node,
		}
	}

	for _, edge := range net.GetEdges() {
		graph.Edges[edge.From.String()] = edge.To.String()
	}

	return graph
}
