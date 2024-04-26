package main

import (
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
	flag.Parse()

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}),
	))

	db, err := bbolt.Open("database.db", 0644, nil)
	if err != nil {
		slog.Error("Cannot open database", "error", err)
		os.Exit(1)
	}

	nodeRepository, err := boltdb.CreateRepository[entity.YggdrasilNode](db)
	if err != nil {
		slog.Error("Cannot create repository for YggdrasilNode", "error", err)
		os.Exit(1)
	}

	client := adminapi.Bind(*yggdrasilSock)
	visitor := StoringVisitor{
		nodesUpdater: nodeRepository,
	}

	walker := netstat.Walker{
		Visitor: visitor,
		Client:  client,
	}

	err = walker.StartFromLocal()
	if err != nil && !errors.Is(err, netstat.ErrStopIteration) {
		slog.Error("Error occurred while network crawling", "error", err)
		os.Exit(1)
	}
}
