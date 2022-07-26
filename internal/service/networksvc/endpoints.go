package networksvc

import (
	"net/http"

	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/gin-gonic/gin"
)

func NewEndpoints(srv IService) *Endpoints {
	return &Endpoints{Service: srv}
}

type Endpoints struct {
	Service IService
}

func (e *Endpoints) Bind(router gin.IRouter) {
	router.GET("", e.GetNetwork)
}

func (e *Endpoints) GetNetwork(ctx *gin.Context) {
	net, err := e.Service.GetNetwork(ctx)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, mapNetworkToDto(net))
}

func mapNetworkToDto(src network.Network) NetworkDto {
	dto := NetworkDto{
		Nodes: make([]NodeDto, len(src.Nodes)),
		Edges: make([]EdgeDto, len(src.Links)),
	}
	for i, node := range src.Nodes {
		dto.Nodes[i] = NodeDto{
			PublicKey:      node.PublicKey.String(),
			AdditionalInfo: node.AdditionalInfo,
		}
	}
	for i, link := range src.Links {
		dto.Edges[i] = EdgeDto{
			From: link.From.String(),
			To:   link.To.String(),
		}
	}
	return dto
}
