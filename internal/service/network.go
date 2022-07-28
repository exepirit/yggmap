package service

import (
	"context"

	"github.com/exepirit/yggmap/internal/domain/network"
)

func NewService(networkRepo network.INetworkRepository) network.INetworkService {
	return &service{networkRepo: networkRepo}
}

type service struct {
	networkRepo network.INetworkRepository
}

func (svc *service) GetNetwork(ctx context.Context) (network.Network, error) {
	return svc.networkRepo.GetCurrent(ctx)
}
