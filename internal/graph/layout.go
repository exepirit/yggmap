package graph

import "github.com/exepirit/yggmap/internal/domain/network"

type Layout interface {
	Assign(graph *network.Graph)
	Run()
}
