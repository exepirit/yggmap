package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/rs/zerolog"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that storing network data in the database.
type StoringVisitor struct {
	logger zerolog.Logger

	repository network.INetworkRepository
	network    *network.Network
}

func (visitor StoringVisitor) VisitNode(node network.Node) bool {
	_ = visitor.network.AddNode(node)
	visitor.logger.Info().
		Str("key", node.PublicKey.String()).
		Msg("New node discovered")
	return len(visitor.network.Nodes) < 10 // TODO: remove limitation after write a tests
}

func (visitor StoringVisitor) VisitLink(from, to network.PublicKey) bool {
	visitor.network.ConnectNodes(from, to)
	return len(visitor.network.Nodes) < 10 // TODO: remove limitation after write a tests
}

func (visitor StoringVisitor) Save(ctx context.Context) error {
	return visitor.repository.Update(ctx, *visitor.network)
}
