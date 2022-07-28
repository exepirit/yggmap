package v1

import (
	"net/http"

	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/gin-gonic/gin"
)

func NewNodeEndpoints(nodeSvc network.INodeService) *NodeEndpoints {
	return &NodeEndpoints{
		nodeSvc: nodeSvc,
	}
}

// NodeEndpoints contains all requests handlers related to network node entity.
type NodeEndpoints struct {
	nodeSvc network.INodeService
}

func (e NodeEndpoints) Bind(router gin.IRouter) {
	router.GET("/active", e.GetActive)
}

func (e NodeEndpoints) GetActive(ctx *gin.Context) {
	activeNodes, err := e.nodeSvc.GetActive(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	nodesDto := make([]NodeDto, len(activeNodes))
	for i, node := range activeNodes {
		nodesDto[i] = NodeDto{
			PublicKey:      node.PublicKey.String(),
			AdditionalInfo: node.AdditionalInfo,
		}
	}

	ctx.JSON(http.StatusOK, nodesDto)
}
