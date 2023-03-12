package main

import (
	"fmt"
	tn "trees_and_graphs"
)

func rightSideView(root *tn.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	q := tn.NewQueue[*tn.TreeNode]()
	q.Push(root)

	for !q.IsEmpty() {
		levelLen := q.Length()
		ans = append(ans, q.PeekLast().Val)

		for i := 0; i < levelLen; i++ {
			node := q.PopLeft()

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}

	return ans
}

func main() {
	root := tn.New(1)
	root.InsertLeft(2)
	root.InsertRight(3)
	root.Right.InsertRight(5)
	root.Right.InsertLeft(4)
	fmt.Println(rightSideView(root)) // [1, 3, 5]
}
