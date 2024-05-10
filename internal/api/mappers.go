package api

import (
	"github.com/exepirit/yggmap/internal/api/dto"
	"github.com/exepirit/yggmap/internal/data/entity"
	"time"
)

func mapYggdrasilNodeToDto(node entity.YggdrasilNode) *dto.YggdrasilNode {
	return &dto.YggdrasilNode{
		Address:   node.Address,
		PublicKey: node.PublicKey.String(),
		LastSeen:  node.LastSeen.UTC().Format(time.RFC3339),
	}
}
