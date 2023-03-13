package main

import (
	"fmt"
	tn "trees_and_graphs"
)

func printAllNodes(root *tn.TreeNode) {
	q := tn.NewQueue[*tn.TreeNode]()
	q.Push(root)

	for !q.IsEmpty() {
		nodeCurLen := q.Length()

		for i := 0; i < nodeCurLen; i++ {
			node := q.PopLeft()

			fmt.Println(node.Val)

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}
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

	printAllNodes(root)
}
