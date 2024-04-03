package sqlite

import (
	"encoding/json"
	"fmt"
	"github.com/exepirit/yggmap/pkg/yggdrasil"
	"strconv"
	"strings"
	"time"
)

type nodeDbo struct {
	PublicKey      []byte    `db:"public_key"`
	Coordinates    string    `db:"coordinates"`
	AdditionalInfo []byte    `db:"additional_info"`
	LastSeen       time.Time `db:"last_seen"`
	IsActive       bool      `db:"is_active"`
}

func mapNodeToAggregate(nodeDbo nodeDbo) (yggdrasil.Node, error) {
	node := yggdrasil.Node{}
	node.PublicKey = nodeDbo.PublicKey

	coordinates := strings.Split(nodeDbo.Coordinates, ",")
	node.Coordinates = make([]int, len(coordinates))
	for i, s := range coordinates {
		var err error
		if node.Coordinates[i], err = strconv.Atoi(s); err != nil {
			return node, fmt.Errorf("cannot parse node coordinates")
		}
	}

	if nodeDbo.AdditionalInfo != nil {
		if err := json.Unmarshal(nodeDbo.AdditionalInfo, &node.AdditionalInfo); err != nil {
			return node, fmt.Errorf("cannot parse node additional info")
		}
	}

	node.LastSeen = nodeDbo.LastSeen
	node.IsActive = nodeDbo.IsActive

	return node, nil
}

func mapAggregateToNode(node yggdrasil.Node) nodeDbo {
	var nodeDbo nodeDbo

	nodeDbo.PublicKey = node.PublicKey
	nodeDbo.AdditionalInfo, _ = json.Marshal(node.AdditionalInfo)
	nodeDbo.LastSeen = node.LastSeen
	nodeDbo.LastSeen = node.LastSeen.UTC().Truncate(time.Second)
	nodeDbo.IsActive = node.IsActive

	for i, c := range node.Coordinates {
		nodeDbo.Coordinates += strconv.Itoa(c)
		if i < len(node.Coordinates)-1 {
			nodeDbo.Coordinates += ","
		}
	}

	return nodeDbo
}

type nodesLinkDbo struct {
	Key1 []byte `db:"key1"`
	Key2 []byte `db:"key2"`
}
