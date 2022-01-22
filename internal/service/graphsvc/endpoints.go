package graphsvc

import (
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewEndpoints(service IService) *Endpoints {
	return &Endpoints{Service: service}
}

type Endpoints struct {
	Service IService
}

func (e Endpoints) Bind(router gin.IRouter) {
	router.GET("", e.GetGraph)
}

func (e Endpoints) GetGraph(ctx *gin.Context) {
	netGraph, err := e.Service.GetGraph(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	result := &GraphDto{
		Nodes: mapNodesToDto(netGraph.Nodes),
		Edges: netGraph.Edges,
	}
	ctx.JSON(http.StatusOK, result)
}

func mapNodesToDto(nodes []network.GraphNode) []NodeDto {
	result := make([]NodeDto, len(nodes))
	for i, n := range nodes {
		result[i] = NodeDto{
			ID:    n.GetID(),
			Label: n.GetID()[len(n.GetID())-8:],
			X:     n.X,
			Y:     n.Y,
		}
	}
	return result
}
