package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func helper(node *tn.TreeNode, curMax, curMin int) int {
	if node == nil {
		return curMax - curMin
	}

	curMax = int(math.Max(float64(curMax), float64(node.Val)))
	curMin = int(math.Min(float64(curMin), float64(node.Val)))

	left := helper(node.Left, curMax, curMin)
	right := helper(node.Right, curMax, curMin)

	return int(math.Max(float64(left), float64(right)))
}

func maxAncestorDiff(root *tn.TreeNode) int {
	if root == nil {
		return 0
	}

	return helper(root, root.Val, root.Val)
}

func main() {
	root := tn.New(8)
	root.InsertLeft(3)
	root.InsertRight(10)
	root.Left.InsertLeft(1)
	root.Left.InsertRight(6)
	root.Left.Right.InsertLeft(4)
	root.Left.Right.InsertRight(7)
	root.Right.InsertRight(14)
	root.Right.Right.InsertLeft(13)

	fmt.Println(maxAncestorDiff(root)) // 7

	root2 := tn.New(1)
	root2.InsertRight(2)
	root2.Right.InsertRight(0)
	root2.Right.Right.InsertLeft(3)

	fmt.Println(maxAncestorDiff(root2)) // 3
}
