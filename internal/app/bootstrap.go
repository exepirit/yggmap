package app

import (
	"context"
	"os"
	"time"

	"github.com/exepirit/yggmap/internal/api"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository"
	"github.com/exepirit/yggmap/internal/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(LoadDbConfig),
	infrastructure.Module,
	repository.Module,
	service.Module,
	api.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	server infrastructure.Server,
	api api.API,
) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822,
	})

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				api.Bind(server.Gin)
				_ = server.Gin.Run(":8000")
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
