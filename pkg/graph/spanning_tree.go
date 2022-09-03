package graph

import "sort"

func MakeSpanningTree[T any](rootNodeData T, nodesData []T, getCoords func(T) []int) Tree[T] {
	root := &TreeNode[T]{
		Data:     rootNodeData,
		Children: make([]*TreeNode[T], 0),
	}

	nodes := make([]TreeNode[T], len(nodesData))
	for i, nodeData := range nodesData {
		nodes[i].Data = nodeData
	}
	sort.Slice(nodes, func(i, j int) bool {
		return len(getCoords(nodes[i].Data)) < len(getCoords(nodes[j].Data))
	})

	for i := range nodes {
		addSTNode(root, getCoords(nodes[i].Data), &nodes[i])
	}

	return Tree[T]{
		Root: root,
	}
}

func addSTNode[T any](root *TreeNode[T], coords []int, node *TreeNode[T]) {
	n := root
	c := coords[0]
	for len(coords) > 0 {
		c = coords[0]
		if len(n.Children) < c+1 {
			n.Children = append(n.Children, make([]*TreeNode[T], c-len(n.Children)+1)...)
		}
		coords = coords[1:]
		if len(coords) > 0 {
			n = n.Children[c]
		}
	}
	n.Children[c] = node
}
