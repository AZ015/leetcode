package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func dfs(node *tn.TreeNode, small, large int) bool {
	if node == nil {
		return true
	}

	if small >= node.Val || node.Val >= large {
		return false
	}

	left := validateBinarySearchTree(node.Left)
	right := validateBinarySearchTree(node.Right)

	return left && right
}

func validateBinarySearchTree(root *tn.TreeNode) bool {
	return dfs(root, math.MaxInt, math.MinInt)
}

func main() {
	root := tn.New(10)
	root.InsertLeft(5)
	root.InsertRight(12)
	root.Left.InsertLeft(2)
	root.Left.InsertRight(8)
	root.Right.InsertRight(23)
	fmt.Println(validateBinarySearchTree(root))
}
