package graph

import (
	"github.com/exepirit/yggmap/internal/domain/network"
	"math/rand"
)

type RandomLayout struct {
	graph *network.Graph
}

func (layout *RandomLayout) Assign(g *network.Graph) {
	layout.graph = g
}

func (layout RandomLayout) Run() {
	for i := range layout.graph.Nodes {
		layout.graph.Nodes[i].X = normalInverse(0, -5)
		layout.graph.Nodes[i].Y = normalInverse(0, -5)
	}
}

func normalInverse(mu float32, sigma float32) float64 {
	return rand.NormFloat64()*float64(sigma) + float64(mu)
}
