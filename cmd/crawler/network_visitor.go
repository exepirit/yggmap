package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"log/slog"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that stores network data in the database.
type StoringVisitor struct {
	network *network.Network
}

func (visitor StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	_ = visitor.network.AddNode(node)
	slog.Info("New node discovered",
		"key", node.PublicKey.String(),
		"ip", node.PublicKey.IPv6Address())
	return len(visitor.network.Nodes) < 10 // TODO: remove limitation after write a tests
}

func (visitor StoringVisitor) VisitLink(from, to yggdrasil.PublicKey) bool {
	visitor.network.ConnectNodes(from, to)
	return len(visitor.network.Nodes) < 10 // TODO: remove limitation after write a tests
}

func (visitor StoringVisitor) Save(ctx context.Context) error {
	// TODO: implement store network information in a database
	panic("not implemented")
}
