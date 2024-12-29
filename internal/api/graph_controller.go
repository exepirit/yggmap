package api

import (
	"github.com/exepirit/yggmap/internal/data/ent"
	"github.com/gofiber/fiber/v2"
)

type GraphController struct {
	Data *ent.Client
}

func (ctrl GraphController) AttachController(router fiber.Router) {
	router.Get("/networkGraph", ctrl.GetGraphData)
}

func (ctrl GraphController) GetGraphData(ctx *fiber.Ctx) error {
	nodes, err := ctrl.Data.YggdrasilNode.
		Query().
		All(ctx.Context())
	if err != nil {
		return err
	}

	data := GraphDataDTO{
		Nodes: make([]graphNodeDTO, 0, len(nodes)),
		Links: make([]graphLinkDTO, 0),
	}
	for _, node := range nodes {
		data.Nodes = append(data.Nodes, graphNodeDTO{ID: node.PublicKey})
		neighbors, err := node.QueryNeighbors().All(ctx.Context())
		if err != nil {
			return err
		}
		for _, neighbor := range neighbors {
			data.Links = append(data.Links, graphLinkDTO{
				Source: node.PublicKey,
				Target: neighbor.PublicKey,
			})
		}
	}

	return ctx.JSON(data)
}

type GraphDataDTO struct {
	Nodes []graphNodeDTO `json:"nodes"`
	Links []graphLinkDTO `json:"links"`
}

type graphNodeDTO struct {
	ID string `json:"id"`
}

type graphLinkDTO struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
