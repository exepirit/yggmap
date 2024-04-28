package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/exepirit/yggmap/internal/api/dto"
)

// Neighbors is the resolver for the neighbors field.
func (r *yggdrasilNodeResolver) Neighbors(ctx context.Context, obj *dto.YggdrasilNode) ([]*dto.YggdrasilNodeLink, error) {
	panic(fmt.Errorf("not implemented: Neighbors - neighbors"))
}

// YggdrasilNode returns YggdrasilNodeResolver implementation.
func (r *Resolver) YggdrasilNode() YggdrasilNodeResolver { return &yggdrasilNodeResolver{r} }

type yggdrasilNodeResolver struct{ *Resolver }
