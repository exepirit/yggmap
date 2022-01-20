package main

import (
	"github.com/exepirit/yggmap/internal/app"
	"go.uber.org/fx"
)

func main() {
	fx.New(app.Module).Run()
}
