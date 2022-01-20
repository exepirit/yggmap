package node_test

import (
	"fmt"
	"testing"

	"github.com/exepirit/yggmap/internal/domain/node"
)

func TestPublicKeyFormatIPv6Address(t *testing.T) {
	testCases := []struct {
		PublicKey   node.PublicKey
		IPv6Address string
	}{
		{
			PublicKey:   node.MustParseKey("b90ad22a1b16412974820e9d163d40e007e3546726f21febec3389b94a5ce228"),
			IPv6Address: "200:8dea:5bab:c9d3:7dad:16fb:e2c5:d385",
		},
		{
			PublicKey:   node.MustParseKey("000000074dac6cf3a02100c646b794655c05534933a34c6ebd80dfe430ccfaca"),
			IPv6Address: "21d:2c94:e4c3:17f7:bfce:6e52:1ae6:a8fe",
		},
		{
			PublicKey:   node.MustParseKey("2493ffffffffec63eb18516ce50a3dc2667e29d49fb8b9bd7b39ffe94a32c882"),
			IPv6Address: "202:db60::9ce0:a73d:7498:d7ae",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Key%d", i), func(t *testing.T) {
			var address string
			address = testCase.PublicKey.IPv6Address()

			if address != testCase.IPv6Address {
				t.Errorf("expected: %q, got: %q", testCase.IPv6Address, address)
			}
		})
	}
}
