package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"time"

	"github.com/exepirit/yggmap/internal/api/dto"
	"github.com/exepirit/yggmap/internal/data"
	"github.com/exepirit/yggmap/internal/data/entity"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, publicKey string) (*dto.YggdrasilNode, error) {
	node, err := r.NodesLoader.Load(ctx, publicKey)
	if err != nil {
		return nil, err
	}

	return &dto.YggdrasilNode{
		Address:   node.Address,
		PublicKey: node.PublicKey.String(),
		LastSeen:  node.LastSeen.UTC().Format(time.RFC3339),
	}, nil
}

// NodesList is the resolver for the nodesList field.
func (r *queryResolver) NodesList(ctx context.Context, previous *string, limit int) (*dto.YggdrasilNodesPage, error) {
	page := &dto.YggdrasilNodesPage{
		Items: make([]*dto.YggdrasilNode, 0, limit),
	}
	err := r.NodesLoader.Provider.Iterate(ctx, func(cursor data.Cursor[entity.YggdrasilNode]) error {
		if previous != nil {
			cursor.Seek(*previous)
		}

		for cursor.Next() != "" && limit > 0 {
			item, err := cursor.Get()
			if err != nil {
				return err
			}

			// TODO: move to mappers
			page.Items = append(page.Items, &dto.YggdrasilNode{
				Address:   item.Address,
				PublicKey: item.PublicKey.String(),
				LastSeen:  item.LastSeen.UTC().Format(time.RFC3339),
			})

			limit--
		}
		return nil
	})
	return page, err
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
