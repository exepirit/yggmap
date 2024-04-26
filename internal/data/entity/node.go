package entity

import (
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"time"
)

type YggdrasilNode struct {
	PublicKey yggdrasil.PublicKey `json:"publicKey"`
	LastSeen  time.Time           `json:"lastSeen"`
}

func (node YggdrasilNode) ID() string {
	return node.PublicKey.String()
}
