package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func maxDepthBinaryTree(node *tn.TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepthBinaryTree(node.Left)
	right := maxDepthBinaryTree(node.Right)

	return int(math.Max(float64(left), float64(right))) + 1
}

func main() {
	root := tn.New(0)
	root.InsertLeft(1)
	root.InsertRight(2)
	root.Left.InsertLeft(3)
	root.Left.InsertRight(4)
	root.Right.InsertRight(5)
	root.Left.Left.InsertLeft(6)
	root.Left.Right.InsertLeft(7)
	root.Left.Right.InsertRight(8)

	fmt.Println(maxDepthBinaryTree(root))
}
