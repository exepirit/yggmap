package main

import (
	"flag"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/exepirit/yggmap/internal/api"
	"github.com/exepirit/yggmap/internal/api/middleware"
	"github.com/exepirit/yggmap/internal/data"
	"github.com/exepirit/yggmap/internal/data/boltdb"
	"github.com/exepirit/yggmap/internal/data/entity"
	"go.etcd.io/bbolt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	dbPath := flag.String("db.path", "database.db", "Database file path")
	flag.Parse()

	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}),
	))

	db, err := bbolt.Open(*dbPath, 0644, nil)
	if err != nil {
		slog.Error("Failed to open the database", "error", err)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{
		NodesLoader: data.Loader[entity.YggdrasilNode]{
			Provider: nodeRepository,
		},
		LinksLoader: data.Loader[entity.NodeLink]{
			Provider: linksRepository,
		},
	}}))
	srv.AroundOperations(middleware.Logging)

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	slog.Info("Listening for client requests", "address", ":"+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Info("HTTP server error", "error", err)
	}
}
