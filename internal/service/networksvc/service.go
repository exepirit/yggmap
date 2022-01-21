package networksvc

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
)

func NewService(networkRepo network.INetworkRepository) IService {
	return &service{networkRepo: networkRepo}
}

type IService interface {
	GetNetwork(ctx context.Context) (*network.Network, error)
}

type service struct {
	networkRepo network.INetworkRepository
}

func (svc *service) GetNetwork(ctx context.Context) (*network.Network, error) {
	return svc.networkRepo.Get(ctx)
}
