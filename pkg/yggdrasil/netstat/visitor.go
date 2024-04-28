package netstat

import "github.com/exepirit/yggmap/pkg/yggdrasil"

// Visitor is an interface that allows users to visit nodes and links in a network.
type Visitor interface {
	// VisitNode visits a node in the network.
	// Returns true if the visitor should continue visiting the graph, false otherwise.
	VisitNode(node yggdrasil.Node) bool

	// VisitLink visits a link between two nodes in the network.
	// Returns true if the visitor should continue visiting the graph deep, false otherwise.
	VisitLink(from, to yggdrasil.PublicKey) bool
}
