package network

import (
	"context"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
)

type INodeRepository interface {
	Get(ctx context.Context, key yggdrasil.PublicKey) (yggdrasil.Node, error)
	GetAll(ctx context.Context) ([]yggdrasil.Node, error)
}

type INetworkRepository interface {
	GetCurrent(ctx context.Context) (Network, error)
	Update(ctx context.Context, network Network) error
}

type ISpanningTreeRepository interface {
	INetworkRepository
}

type INetworkService interface {
	GetNetwork(ctx context.Context) (Network, error)
	GetSpanningTree(ctx context.Context) (Network, error)
}

type INodeService interface {
	GetActive(ctx context.Context) ([]yggdrasil.Node, error)
}
