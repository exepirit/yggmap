package network

import "context"

type INetworkRepository interface {
	Update(ctx context.Context, net *Network) error
	Get(ctx context.Context) (*Network, error)
}

type INodeRepository interface {
	Get(ctx context.Context, key PublicKey) (*Node, error)
	GetAll(ctx context.Context) ([]*Node, error)
	UpdateAll(ctx context.Context, nodes []*Node) error
}
