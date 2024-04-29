package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/exepirit/yggmap/internal/api/dto"
)

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, skip int, limit int) (*dto.YggdrasilNodesPage, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }