package service

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"

	"github.com/exepirit/yggmap/internal/domain/network"
)

// NewNodeService creates new instance of network.INodeService.
func NewNodeService(nodes network.INodeRepository) network.INodeService {
	return &nodeService{
		nodes: nodes,
	}
}

type nodeService struct {
	nodes network.INodeRepository
}

func (svc nodeService) GetActive(ctx context.Context) ([]yggdrasil.Node, error) {
	allNodes, err := svc.nodes.GetAll(ctx)
	if err != nil {
		return allNodes, fmt.Errorf("cannot get all nodes list: %w", err)
	}

	activeNodes := make([]yggdrasil.Node, 0)
	for _, node := range allNodes {
		if node.IsActive {
			activeNodes = append(activeNodes, node)
		}
	}

	return activeNodes, nil
}
