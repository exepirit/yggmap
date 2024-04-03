package main

import (
	"context"
	"flag"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"github.com/exepirit/yggmap/pkg/yggdrasil/netstat"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	yggdrasilSock := flag.String("socket", "unix:///var/run/yggdrasil/yggdrasil.sock", "Yggdrasil API socket")
	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	client := adminapi.Bind(*yggdrasilSock)
	visitor := StoringVisitor{
		logger:  log.Logger,
		network: &network.Network{},
	}

	err := netstat.WalkNetwork(context.Background(), client, visitor)
	if err != nil {
		log.Fatal().Err(err).Msg("Error occurred while network crawling")
	}

	if err = visitor.Save(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Cannot save network")
	}
}
