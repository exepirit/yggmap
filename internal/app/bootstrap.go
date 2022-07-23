package app

import (
	"context"
	"os"
	"time"

	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository"
	"github.com/exepirit/yggmap/internal/service/graphsvc"
	"github.com/exepirit/yggmap/internal/service/networksvc"
	"github.com/exepirit/yggmap/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(LoadDbConfig),
	infrastructure.Module,
	repository.Module,
	networksvc.Module,
	graphsvc.Module,
	fx.Provide(NewRoutes),
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	server infrastructure.Server,
	routes server.Bindable,
) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				routes.Bind(server.Gin)
				_ = server.Gin.Run(":8000")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
