package main

import (
	"github.com/exepirit/yggmap/internal/app"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.Module,
		fx.WithLogger(infrastructure.NewFxLogger),
	).Run()
}
