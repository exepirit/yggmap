package v1

import (
	"github.com/exepirit/yggmap/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewNetworkEndpoints),
	fx.Provide(NewNodeEndpoints),
	fx.Provide(NewAPI),
)

type API server.Bindable

func NewAPI(
	network *NetworkEndpoints,
	node *NodeEndpoints,
) API {
	return server.Union(
		server.Route("/network", network),
		server.Route("/node", node),
	)
}
