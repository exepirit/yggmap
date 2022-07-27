package api

import (
	v1 "github.com/exepirit/yggmap/internal/api/v1"
	"github.com/exepirit/yggmap/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	v1.Module,
	fx.Provide(NewAPI),
)

type API server.Bindable

func NewAPI(apiv1 v1.API) API {
	return server.Route("/api",
		server.Route("/v1", apiv1),
	)
}
