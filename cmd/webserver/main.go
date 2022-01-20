package main

import (
	"github.com/exepirit/yggmap/internal/app"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		app.Module,
		fx.WithLogger(func(logger infrastructure.Logger) fxevent.Logger {
			return logger.GetFxLogger()
		}),
	).Run()
}
