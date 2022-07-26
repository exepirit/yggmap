package app

import (
	"github.com/exepirit/yggmap/internal/service/networksvc"
	"github.com/exepirit/yggmap/pkg/server"
)

func NewRoutes(
	networkEndpoints *networksvc.Endpoints,
) server.Bindable {
	return server.Route("/api",
		server.Route("/network", networkEndpoints))
}
