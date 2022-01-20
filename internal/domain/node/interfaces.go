package node

import "context"

type Repository interface {
	Put(ctx context.Context, node *Node) error
	PutOrUpdate(ctx context.Context, node *Node) error
	Get(ctx context.Context, key PublicKey) (*Node, error)
	GetAll(ctx context.Context) ([]*Node, error)
}
