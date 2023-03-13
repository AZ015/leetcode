package main

import (
	"fmt"
	tn "trees_and_graphs"
)

func rangeSumOfBst(root *tn.TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Val > low && root.Val < high {
		sum += root.Val
	}

	if low < root.Val {
		sum += rangeSumOfBst(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumOfBst(root.Right, low, high)

	}

	return sum
}

func main() {
	root := tn.New(10)
	root.InsertLeft(5)
	root.InsertRight(15)
	root.Left.InsertLeft(3)
	root.Left.InsertRight(7)
	root.Right.InsertRight(18)
	fmt.Println(rangeSumOfBst(root, 7, 15)) // 25
}
