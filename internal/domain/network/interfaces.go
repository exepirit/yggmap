package network

import "context"

type IRepository interface {
	Update(ctx context.Context, net *Network) error
	Get(ctx context.Context) (*Network, error)
}
