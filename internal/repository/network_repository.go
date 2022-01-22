package repository

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
)

func NewNetworkRepository(nodes network.INodeRepository) network.INetworkRepository {
	return &NetworkRepositoryMongoDb{
		nodes: nodes,
	}
}

type NetworkRepositoryMongoDb struct {
	nodes network.INodeRepository
}

func (repo *NetworkRepositoryMongoDb) Update(ctx context.Context, net *network.Network) error {
	return repo.nodes.UpdateAll(ctx, net.Nodes)
}

func (repo *NetworkRepositoryMongoDb) Get(ctx context.Context) (*network.Network, error) {
	nodes, err := repo.nodes.GetAll(ctx)
	return &network.Network{
		Nodes: nodes,
	}, err
}
