package main

import (
	"context"
	"errors"
	"flag"
	"github.com/exepirit/yggmap/internal/data/db"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"github.com/exepirit/yggmap/pkg/yggdrasil/netstat"
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

	dbClient, err := db.ConnectSqlite(context.Background(), db.SqliteConfig{Path: "database.sqlite"})
	if err != nil {
		slog.Error("Failed to create database connection", "error", err)
		os.Exit(1)
	}
	defer dbClient.Close()

	visitor := &StoringVisitor{
		Client: dbClient,
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
