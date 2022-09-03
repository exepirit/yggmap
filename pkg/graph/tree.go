package graph

type Tree[T any] struct {
	Root *TreeNode[T]
}

type TreeNode[T any] struct {
	Data     T
	Children []*TreeNode[T]
}

func WalkDepth[T any](tree Tree[T], callback func(parent *T, node T)) {
	callback(nil, tree.Root.Data)
	nodeWalkDepth(tree.Root, callback)
}

func nodeWalkDepth[T any](node *TreeNode[T], callback func(parent *T, node T)) {
	for _, child := range node.Children {
		if child != nil {
			callback(&node.Data, child.Data)
			nodeWalkDepth(child, callback)
		}
	}
}
