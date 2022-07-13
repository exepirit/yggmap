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

	ctx.JSON(http.StatusOK, mapGraphToDto(netGraph))
}

func mapGraphToDto(graph *network.Graph) GraphDto {
	dto := GraphDto{
		Nodes: make([]NodeDto, len(graph.Nodes)),
		Edges: make([]EdgeDto, len(graph.Edges)),
	}
	for i, n := range graph.Nodes {
		dto.Nodes[i] = NodeDto{
			ID:    n.GetID(),
			Label: n.GetID()[len(n.GetID())-8:],
			X:     n.X,
			Y:     n.Y,
		}
	}
	for i, e := range graph.Edges {
		dto.Edges[i] = EdgeDto{
			From: e.From.String(),
			To:   e.To.String(),
		}
	}
	return dto
}
