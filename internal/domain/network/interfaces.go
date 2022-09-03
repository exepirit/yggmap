package network

import "context"

type INodeRepository interface {
	Get(ctx context.Context, key PublicKey) (Node, error)
	GetAll(ctx context.Context) ([]Node, error)
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
	GetActive(ctx context.Context) ([]Node, error)
}
