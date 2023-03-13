package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func minDepth(node *tn.TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	min := math.MaxInt

	if node.Left != nil {
		min = int(math.Min(float64(minDepth(node.Left)), float64(min)))
	}

	if node.Right != nil {
		min = int(math.Min(float64(minDepth(node.Right)), float64(min)))
	}

	return min + 1
}

func main() {
	root := tn.New(3)
	root.InsertLeft(9)
	root.InsertRight(20)
	root.Right.InsertRight(7)
	root.Right.InsertLeft(15)

	fmt.Println(minDepth(root)) // 2

	root2 := tn.New(2)
	root2.InsertRight(3)
	root2.Right.InsertRight(4)
	root2.Right.Right.InsertRight(5)
	root2.Right.Right.Right.InsertRight(6)

	fmt.Println(minDepth(root2))
}
