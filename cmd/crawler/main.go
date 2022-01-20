package main

import (
	"context"
	"github.com/exepirit/yggmap/internal/crawl"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	dbConf := infrastructure.DatabaseConfig{
		URI:  os.Getenv("MONGODB_URI"),
		Name: os.Getenv("MONGODB_NAME"),
	}
	database, err := infrastructure.NewDatabase(dbConf)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to connect to database")
	}

	netRepo := repository.NewNetworkRepository(database)

	client := adminapi.Bind("unix:///var/run/yggdrasil.sock")
	crawler := crawl.NetworkCrawler{Client: client}

	net, err := crawler.GetNetwork(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to crawl whole network")
	}

	log.Info().Msg("Network successfully crawled!")
	log.Info().Msgf("Collected %d nodes with %d connections", len(net.Nodes), len(net.Edges))

	if err = netRepo.Update(context.Background(), net); err != nil {
		log.Error().Err(err).Msg("Failed to update network in database")
	}
}
