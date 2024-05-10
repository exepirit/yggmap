package filter

import (
	"github.com/exepirit/yggmap/internal/data/entity"
	"strings"
)

func NodeAddressContains(substr string) Predicate[entity.YggdrasilNode] {
	substr = strings.ToLower(substr)
	return func(node entity.YggdrasilNode) bool {
		return strings.Contains(node.Address, substr)
	}
}

func NodeKeyContainsStr(substr string) Predicate[entity.YggdrasilNode] {
	substr = strings.ToLower(substr)
	return func(node entity.YggdrasilNode) bool {
		return strings.Contains(node.PublicKey.String(), substr)
	}
}
