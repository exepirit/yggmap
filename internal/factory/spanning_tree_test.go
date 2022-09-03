package factory_test

import (
	"github.com/exepirit/yggmap/internal/factory"
	"testing"

	"github.com/exepirit/yggmap/internal/domain/network"
)

func TestMakeSpanningTree(t *testing.T) {
	nodes := []network.Node{
		{
			PublicKey:   network.MustParseKey("f5c3c93e8206c93847c3af418e18e696f5c3c93e8206c93847c3af418e18e696"),
			Coordinates: []int{},
		},
		{
			PublicKey:   network.MustParseKey("14246b5cce1fef3da0adea26a72a270a14246b5cce1fef3da0adea26a72a270a"),
			Coordinates: []int{0},
		},
		{
			PublicKey:   network.MustParseKey("22308363d33e3e594a2850f29241d9b222308363d33e3e594a2850f29241d9b2"),
			Coordinates: []int{0, 1},
		},
		{
			PublicKey:   network.MustParseKey("8b90dccb470ab8aa9c15510798c208278b90dccb470ab8aa9c15510798c20827"),
			Coordinates: []int{1},
		},
		{
			PublicKey:   network.MustParseKey("6208815a0a3659b8b5139b921bbea3d26208815a0a3659b8b5139b921bbea3d2"),
			Coordinates: []int{0, 1, 2},
		},
		{
			PublicKey:   network.MustParseKey("135095633968b58a5e781a7f0e28cae3135095633968b58a5e781a7f0e28cae3"),
			Coordinates: []int{1, 3},
		},
		{
			PublicKey:   network.MustParseKey("49d47371f96fcb2f1da9f7c74c3122e049d47371f96fcb2f1da9f7c74c3122e0"),
			Coordinates: []int{1, 2},
		},
	}

	spanningTree := factory.MakeSpanningTree(nodes)

	checkLink := func(k1, k2 string) {
		key1, key2 := network.MustParseKey(k1), network.MustParseKey(k2)
		for _, link := range spanningTree.Links {
			if link.From.Equal(key1) && link.To.Equal(key2) {
				return
			}
			if link.To.Equal(key1) && link.From.Equal(key2) {
				return
			}
		}
		t.Logf("link between two nodes not found\n\tkey1 = %s\n\tkey2 = %s", k1, k2)
		t.Fail()
	}
	checkLink(
		"f5c3c93e8206c93847c3af418e18e696f5c3c93e8206c93847c3af418e18e696",
		"14246b5cce1fef3da0adea26a72a270a14246b5cce1fef3da0adea26a72a270a")
	checkLink(
		"f5c3c93e8206c93847c3af418e18e696f5c3c93e8206c93847c3af418e18e696",
		"8b90dccb470ab8aa9c15510798c208278b90dccb470ab8aa9c15510798c20827")
	checkLink(
		"14246b5cce1fef3da0adea26a72a270a14246b5cce1fef3da0adea26a72a270a",
		"22308363d33e3e594a2850f29241d9b222308363d33e3e594a2850f29241d9b2")
	checkLink(
		"22308363d33e3e594a2850f29241d9b222308363d33e3e594a2850f29241d9b2",
		"6208815a0a3659b8b5139b921bbea3d26208815a0a3659b8b5139b921bbea3d2")
	checkLink(
		"8b90dccb470ab8aa9c15510798c208278b90dccb470ab8aa9c15510798c20827",
		"135095633968b58a5e781a7f0e28cae3135095633968b58a5e781a7f0e28cae3")
	checkLink(
		"8b90dccb470ab8aa9c15510798c208278b90dccb470ab8aa9c15510798c20827",
		"49d47371f96fcb2f1da9f7c74c3122e049d47371f96fcb2f1da9f7c74c3122e0")
}
