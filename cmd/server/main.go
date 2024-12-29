package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/exepirit/yggmap/internal/api"
	"github.com/exepirit/yggmap/internal/data/db"
	web "github.com/exepirit/yggmap/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	flag.Parse()
	
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelInfo,
		}),
	)
	slog.SetDefault(logger)

	dbClient, err := db.ConnectSqlite(context.Background(), db.SqliteConfig{Path: "database.sqlite"})
	if err != nil {
		slog.Error("Failed to create database connection", "error", err)
		os.Exit(1)
	}


	app := fiber.New()
	app.Use(slogfiber.New(logger))
	app.Use(recover.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	apiRouter := app.Group("/api")
	graphController := api.GraphController{Data: dbClient}
	graphController.AttachController(apiRouter)

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(web.Static),
		PathPrefix: "dist",
		Browse: true,
		Index: "index.html",
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Listen(":"+port)
}
