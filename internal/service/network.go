package service

import (
	"context"
	"github.com/exepirit/yggmap/internal/factory"

	"github.com/exepirit/yggmap/internal/domain/network"
)

func NewNetworkService(networkRepo network.INetworkRepository) network.INetworkService {
	return &service{networkRepo: networkRepo}
}

type service struct {
	networkRepo network.INetworkRepository
}

func (svc *service) GetNetwork(ctx context.Context) (network.Network, error) {
	return svc.networkRepo.GetCurrent(ctx)
}

func (svc *service) GetSpanningTree(ctx context.Context) (network.Network, error) {
	net, err := svc.GetNetwork(ctx)
	if err != nil {
		return network.Network{}, err
	}

	return factory.MakeSpanningTree(net.Nodes), nil
}
