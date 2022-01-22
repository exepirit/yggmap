package graphsvc

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/internal/factory"
	"github.com/exepirit/yggmap/internal/graph"
)

type IService interface {
	GetGraph(ctx context.Context) (*network.Graph, error)
}

func NewService(networkRepo network.INetworkRepository) IService {
	return &service{networkRepo: networkRepo}
}

type service struct {
	networkRepo network.INetworkRepository
}

func (svc service) GetGraph(ctx context.Context) (*network.Graph, error) {
	net, err := svc.networkRepo.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query network: %w", err)
	}

	netGraph := factory.MakeGraphForNetwork(net)
	layout := graph.RandomLayout{}
	layout.Assign(netGraph)
	layout.Run()

	return netGraph, nil
}
