package repository

import (
	"github.com/exepirit/yggmap/internal/domain/network"
	"github.com/exepirit/yggmap/internal/infrastructure"
	"github.com/exepirit/yggmap/internal/repository/sqlite"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewNodeRepository),
	fx.Provide(NewNetworkRepository),
)

func NewNodeRepository(database infrastructure.Database) network.INodeRepository {
	return sqlite.NewNodeRepository(database.SQL)
}

func NewNetworkRepository(database infrastructure.Database) network.INetworkRepository {
	return sqlite.NewNetworkRepository(database.SQL)
}
