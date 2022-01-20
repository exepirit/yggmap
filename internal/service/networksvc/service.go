package networksvc

import (
	"context"
	"github.com/exepirit/yggmap/internal/domain/network"
)

func NewService(networkRepo network.IRepository) IService {
	return &service{networkRepo: networkRepo}
}

type IService interface {
	GetNetwork(ctx context.Context) (*network.Network, error)
}

type service struct {
	networkRepo network.IRepository
}

func (svc *service) GetNetwork(ctx context.Context) (*network.Network, error) {
	return svc.networkRepo.Get(ctx)
}
