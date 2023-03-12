package main

import (
	"fmt"
	tn "trees_and_graphs"
)

func deepestLeavesSum(root *tn.TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	q := tn.NewQueue[*tn.TreeNode]()
	q.Push(root)

	for !q.IsEmpty() {
		levelLen := q.Length()
		levelSum := 0

		for i := 0; i < levelLen; i++ {
			node := q.PopLeft()

			if node.Left == nil && node.Right == nil {
				levelSum += node.Val
			}

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		sum = levelSum
	}

	return sum
}

func deepestLeavesSumOptimize(root *tn.TreeNode) int {
	if root == nil {
		return 0
	}

	nextLevel := tn.NewQueue[*tn.TreeNode]()
	currentLevel := tn.NewQueue[*tn.TreeNode]()
	nextLevel.Push(root)

	for !nextLevel.IsEmpty() {
		currentLevel.Copy(nextLevel)
		nextLevel.Clear()

		for _, node := range currentLevel.Queue {
			if node.Left != nil {
				nextLevel.Push(node.Left)
			}

			if node.Right != nil {
				nextLevel.Push(node.Right)
			}
		}
	}

	sum := 0
	for !currentLevel.IsEmpty() {
		node := currentLevel.PopLeft()
		sum += node.Val
	}

	return sum
}

func main() {
	root := tn.New(1)
	root.InsertLeft(2)
	root.InsertRight(3)
	root.Left.InsertLeft(4)
	root.Left.InsertRight(5)
	root.Right.InsertRight(6)
	root.Left.Left.InsertLeft(7)
	root.Right.Right.InsertRight(8)
	fmt.Println(deepestLeavesSum(root))         // 15
	fmt.Println(deepestLeavesSumOptimize(root)) // 15
}
