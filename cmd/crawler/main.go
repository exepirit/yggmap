package main

import (
	"context"
	"errors"
	"flag"
	"github.com/exepirit/yggmap/internal/data/boltdb"
	"github.com/exepirit/yggmap/internal/data/entity"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"github.com/exepirit/yggmap/pkg/yggdrasil/netstat"
	"go.etcd.io/bbolt"
	"log/slog"
	"os"
)

func main() {
	yggdrasilSock := flag.String("socket", "unix:///var/run/yggdrasil/yggdrasil.sock", "Yggdrasil API socket")
	dbPath := flag.String("db.path", "database.db", "Database file path")
	flag.Parse()

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		}),
	))

	db, err := bbolt.Open(*dbPath, 0644, nil)
	if err != nil {
		slog.Error("Failed to open the database", "path", *dbPath, "error", err)
		os.Exit(1)
	}

	nodeRepository, err := boltdb.CreateRepository[entity.YggdrasilNode](db)
	if err != nil {
		slog.Error("Failed to create the YggdrasilNode repository", "error", err)
		os.Exit(1)
	}
	linksRepository, err := boltdb.CreateRepository[entity.NodeLink](db)
	if err != nil {
		slog.Error("Failed to create the NodeLink repository", "error", err)
		os.Exit(1)
	}
	snapshotRepository, err := boltdb.CreateRepository[entity.SnapshotMeta](db)
	if err != nil {
		slog.Error("Failed to create the SnapshotMeta repository", "error", err)
		os.Exit(1)
	}
	visitor := &StoringVisitor{
		nodesUpdater:    nodeRepository,
		linksUpdater:    linksRepository,
		snapshotUpdater: snapshotRepository,
	}

	client := adminapi.Bind(*yggdrasilSock)
	walker := netstat.Walker{
		Visitor: visitor,
		Client:  client,
	}

	err = walker.StartFromLocal()
	if err != nil && !errors.Is(err, netstat.ErrStopIteration) {
		slog.Error("Failed to start the network crawling", "socket", *yggdrasilSock, "error", err)
		os.Exit(1)
	}

	err = visitor.Save(context.Background())
	if err != nil {
		slog.Error("Failed to save the network graph in the database",
			"databasePath", *dbPath, "error", err)
		os.Exit(1)
	}
}
