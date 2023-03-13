package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func helper(node *tn.TreeNode, diameter int) int {
	if node == nil {
		return 0
	}

	left := helper(node.Left, diameter)
	right := helper(node.Right, diameter)
	diameter = int(math.Max(float64(diameter), float64(left+right)))

	return int(math.Max(float64(left), float64(right))) + 1
}

func diameterOfBinaryTree(root *tn.TreeNode) int {
	return helper(root, 0)
}

func main() {
	root := tn.New(2)
	root.InsertLeft(2)
	root.InsertRight(3)
	root.Left.InsertLeft(4)
	root.Left.InsertRight(5)

	fmt.Println(diameterOfBinaryTree(root)) // 3

	root2 := tn.New(1)
	root2.InsertLeft(2)

	fmt.Println(diameterOfBinaryTree(root2)) // 2
}
