package main

import (
	tn "trees_and_graphs"
)

func insertIntoBST(root *tn.TreeNode, val int) *tn.TreeNode {
	if root == nil {
		return tn.New(val)
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}

	return root
}

func main() {
	root := tn.New(4)
	root.InsertLeft(2)
	root.InsertRight(7)
	root.Left.InsertLeft(1)
	root.Left.InsertRight(3)
	tn.PreorderDfs(insertIntoBST(root, 5)) // [4, 2, 7, 1, 3, 5]
}
