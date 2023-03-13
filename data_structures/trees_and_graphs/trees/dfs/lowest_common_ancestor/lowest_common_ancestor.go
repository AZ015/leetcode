package main

import (
	"fmt"
	tn "trees_and_graphs"
)

func lowestCommonAncestor(root, p, q *tn.TreeNode) *tn.TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func main() {
	root := tn.New(3)
	root.InsertLeft(5)
	root.InsertRight(1)
	root.Left.InsertLeft(6)
	root.Left.InsertRight(2)
	root.Right.InsertLeft(7)
	root.Right.InsertRight(4)
	root.Right.InsertLeft(0)
	root.Right.InsertRight(8)

	p := tn.New(6)
	q := tn.New(2)

	fmt.Println(lowestCommonAncestor(root, p, q))
}
