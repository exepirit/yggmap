package main

import (
	"context"
	"fmt"
	"github.com/exepirit/yggmap/internal/data/ent"
	"iter"
	"log/slog"
)

func clusterGraph(ctx context.Context, client *ent.YggdrasilNodeClient) error {
	nodes := make(map[int]int) // store node and color
	counter := 0
	for node := range iterNodes(ctx, client, 50) {
		nodes[node.ID] = counter
		counter++
	}

	iterations := 0
	for iterations < 2 {
		for id := range nodes {
			neighborsIds, err := getNeighborsIds(ctx, client, id)
			if err != nil {
				return err
			}
			neighborColors := make(map[int]int) // store color and count of neighbors with this color
			for _, neighborId := range neighborsIds {
				neighborColor, ok := nodes[neighborId]
				if !ok {
					continue
				}
				neighborColors[neighborColor]++
			}

			if len(neighborColors) == 0 {
				continue
			}

			var newColor, maxClusterSize int
			for color, clusterSize := range neighborColors {
				if clusterSize > maxClusterSize {
					newColor, maxClusterSize = color, clusterSize
				}
			}
			nodes[id] = newColor
		}
		iterations++
	}

	for id, color := range nodes {
		_, err := client.UpdateOneID(id).SetCluster(color).Save(ctx)
		if err != nil {
			return err
		}
		slog.Debug("Set node color", "nodeId", id, "color", color)
	}

	return nil
}

func iterNodes(ctx context.Context, client *ent.YggdrasilNodeClient, chunkSize int) iter.Seq[*ent.YggdrasilNode] {
	return func(yield func(*ent.YggdrasilNode) bool) {
		offset := 0

		for {
			nodes, err := client.Query().Offset(offset).Limit(chunkSize).All(ctx) // TODO: handle error
			if err != nil {
				slog.Error("Failed to query nodes", "error", err)
			}

			for _, node := range nodes {
				if !yield(node) {
					return
				}
			}

			offset += len(nodes)
			if len(nodes) < chunkSize {
				return
			}
		}
	}
}

func getNeighborsIds(ctx context.Context, client *ent.YggdrasilNodeClient, id int) ([]int, error) {
	node, err := client.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query node %d: %w", id, err)
	}

	neighbors, err := node.QueryNeighbors().IDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to query node %d neighbors: %w", id, err)
	}
	return neighbors, nil
}
