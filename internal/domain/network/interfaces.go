package network

import "context"

type INetworkRepository interface {
	Update(ctx context.Context, net *Network) error
	Get(ctx context.Context) (*Network, error)
}

type INodeRepository interface {
	Put(ctx context.Context, node *Node) error
	PutOrUpdate(ctx context.Context, node *Node) error
	Get(ctx context.Context, key PublicKey) (*Node, error)
	GetAll(ctx context.Context) ([]*Node, error)
}
