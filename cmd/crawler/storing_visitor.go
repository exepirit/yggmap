package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/data"
	"github.com/exepirit/yggmap/internal/data/entity"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"log/slog"
	"time"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that stores network data in the database.
type StoringVisitor struct {
	foundNodes     []entity.YggdrasilNode
	nodesAdjacency []entity.NodeLink

	nodesUpdater data.Updater[entity.YggdrasilNode]
	linksUpdater data.Updater[entity.NodeLink]
}

func (visitor *StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	slog.Info("Found a new node", "key", node.PublicKey.String(), "address", node.PublicKey.IPv6Address())
	visitor.foundNodes = append(visitor.foundNodes, entity.YggdrasilNode{
		Address:   node.Address(),
		PublicKey: node.PublicKey,
		LastSeen:  time.Now(),
	})
	return true
}

func (visitor *StoringVisitor) VisitLink(from, to yggdrasil.PublicKey) bool {
	visitor.nodesAdjacency = append(visitor.nodesAdjacency, entity.NodeLink{
		Out:      from,
		In:       to,
		LastSeen: time.Now(),
	})
	return true
}

func (visitor *StoringVisitor) Save(ctx context.Context) error {
	err := visitor.nodesUpdater.PutBatch(ctx, visitor.foundNodes...)
	if err != nil {
		return err
	}
	slog.Info("Nodes stored in the database", "count", len(visitor.foundNodes))

	return visitor.linksUpdater.PutBatch(ctx, visitor.nodesAdjacency...)
}
