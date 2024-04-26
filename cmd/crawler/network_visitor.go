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
	nodesUpdater data.Updater[entity.YggdrasilNode]
}

func (visitor StoringVisitor) VisitNode(node yggdrasil.Node) bool {
	ent := entity.YggdrasilNode{
		PublicKey: node.PublicKey,
		LastSeen:  time.Now(),
	}
	slog.Info("New node discovered", "id", ent.ID())

	err := visitor.nodesUpdater.PutBatch(context.Background(), ent)
	if err != nil {
		slog.Error("Failed to store node", "error", err)
	}

	return true
}

func (visitor StoringVisitor) VisitLink(from, to yggdrasil.PublicKey) bool {
	return true
}
