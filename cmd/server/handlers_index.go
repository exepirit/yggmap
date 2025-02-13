package main

import (
	"github.com/a-h/templ"
	"github.com/exepirit/yggmap/web/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func (h *Handlers) HandleIndex(ctx *fiber.Ctx) error {
	nodes, err := h.database.YggdrasilNode.
		Query().
		All(ctx.Context())
	if err != nil {
		return err
	}

	data := app.GraphData{
		Nodes: make([]app.NodeData, 0, len(nodes)),
		Links: make([]app.LinkData, 0),
	}
	for _, node := range nodes {
		data.Nodes = append(data.Nodes, app.NodeData{
			ID:    node.PublicKey,
			Group: node.Cluster,
		})
		neighbors, err := node.QueryNeighbors().All(ctx.Context())
		if err != nil {
			return err
		}
		for _, neighbor := range neighbors {
			data.Links = append(data.Links, app.LinkData{
				Source: node.PublicKey,
				Target: neighbor.PublicKey,
			})
		}
	}

	requestHandler := templ.Handler(app.Index(data))
	return adaptor.HTTPHandler(requestHandler)(ctx)
}
