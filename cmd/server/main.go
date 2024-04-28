package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/exepirit/yggmap/internal/api"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	slog.SetDefault(slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}),
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(api.NewExecutableSchema(api.Config{Resolvers: &api.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	slog.Info("Listening for client requests", "address", ":"+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		slog.Info("HTTP server error", "error", err)
	}
}
