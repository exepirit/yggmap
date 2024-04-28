package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/data"
	"github.com/exepirit/yggmap/internal/data/entity"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

// StoringVisitor a crawl.NetworkVisitor implementation, that stores network data in the database.
type StoringVisitor struct {
	foundNodes     []entity.YggdrasilNode
	nodesAdjacency []entity.NodeLink

	nodesUpdater    data.Updater[entity.YggdrasilNode]
	linksUpdater    data.Updater[entity.NodeLink]
	snapshotUpdater data.Updater[entity.SnapshotMeta]
}

func (visitor *StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	slog.Info("Found the network node", "key", node.PublicKey.String(), "address", node.PublicKey.IPv6Address())
	visitor.foundNodes = append(visitor.foundNodes, entity.YggdrasilNode{
		Address:   node.Address(),
		PublicKey: node.PublicKey,
		LastSeen:  time.Now(),
	})
	return len(visitor.foundNodes) < 20
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

	err = visitor.linksUpdater.PutBatch(ctx, visitor.nodesAdjacency...)
	if err != nil {
		return err
	}
	slog.Info("Links stored in the database", "count", len(visitor.nodesAdjacency))

	snapshot := entity.SnapshotMeta{
		Identifier: uuid.New(),
		CapturedAt: time.Now(),
		Nodes:      make([]string, 0, len(visitor.foundNodes)),
	}
	for _, node := range visitor.foundNodes {
		snapshot.Nodes = append(snapshot.Nodes, node.ID())
	}

	return visitor.snapshotUpdater.PutBatch(ctx, snapshot)
}
