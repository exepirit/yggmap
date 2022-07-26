package main

import (
	"context"
	"os"
	"time"

	"github.com/exepirit/yggmap/internal/crawl"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository"
	"github.com/exepirit/yggmap/pkg/adminapi"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	dbConf := infrastructure.DatabaseConfig{
		Type:             os.Getenv("DB_TYPE"),
		ConnectionString: os.Getenv("DB_CONNECTIONSTRING"),
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
	log.Info().Msgf("Collected %d nodes", len(net.Nodes))

	if err = netRepo.Update(context.Background(), net); err != nil {
		log.Error().Err(err).Msg("Failed to update network in database")
	}
}
