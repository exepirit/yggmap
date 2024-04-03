package main

import (
	"context"
	"flag"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"github.com/exepirit/yggmap/pkg/yggdrasil/netstat"
	"os"
	"time"

	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	dbType := flag.String("dbType", "sqlite3", "database type")
	dbConnectionString := flag.String("dbConnStr", "yggdrasil_network.db", "database connection url")
	yggdrasilSock := flag.String("socket", "unix:///var/run/yggdrasil.sock", "Yggdrasil API socket")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	dbConf := infrastructure.DatabaseConfig{
		Type:             *dbType,
		ConnectionString: *dbConnectionString,
	}
	database, err := infrastructure.NewDatabase(dbConf)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to connect to database")
	}
	netRepo := repository.NewNetworkRepository(database)

	client := adminapi.Bind(*yggdrasilSock)
	visitor := StoringVisitor{
		logger:     log.Logger,
		repository: netRepo,
		network:    &network.Network{},
	}

	err = netstat.WalkNetwork(context.Background(), client, visitor)
	if err != nil {
		log.Fatal().Err(err).Msg("Error occurred while network crawling")
	}

	if err = visitor.Save(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Cannot save network")
	}
}
