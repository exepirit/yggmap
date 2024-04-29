package api

import (
	"github.com/exepirit/yggmap/internal/data"
	"github.com/exepirit/yggmap/internal/data/entity"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	NodesLoader data.Loader[entity.YggdrasilNode]
	LinksLoader data.Loader[entity.NodeLink]
}
