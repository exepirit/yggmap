package main

import (
	"context"
	"flag"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/pkg/yggdrasil/adminapi"
	"github.com/exepirit/yggmap/pkg/yggdrasil/netstat"
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

	client := adminapi.Bind(*yggdrasilSock)
	visitor := StoringVisitor{
		network: &network.Network{},
	}

	walker := netstat.Walker{
		Visitor: visitor,
		Client:  client,
	}

	err := walker.StartFromLocal()
	if err != nil {
		slog.Error("Error occurred while network crawling", "error", err)
		os.Exit(1)
	}

	if err = visitor.Save(context.Background()); err != nil {
		slog.Error("Cannot store network in database", "error", err)
		os.Exit(1)
	}
}
