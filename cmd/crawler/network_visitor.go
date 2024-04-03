package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/rs/zerolog"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that stores network data in the database.
type StoringVisitor struct {
	logger  zerolog.Logger
	network *network.Network
}

func (visitor StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	_ = visitor.network.AddNode(node)
	visitor.logger.Info().
		Str("key", node.PublicKey.String()).
		Msg("New node discovered")
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
