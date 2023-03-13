package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func largestValInEachTreeRow(root *tn.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := make([]int, 0)
	q := tn.NewQueue[*tn.TreeNode]()
	q.Push(root)

	for !q.IsEmpty() {
		levelLen := q.Length()
		maxLevelVal := math.MinInt

		for i := 0; i < levelLen; i++ {
			node := q.PopLeft()
			maxLevelVal = int(math.Max(float64(maxLevelVal), float64(node.Val)))

			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		ans = append(ans, maxLevelVal)
	}

	return ans
}

func main() {
	root := tn.New(1)
	root.InsertLeft(3)
	root.InsertRight(2)
	root.Left.InsertLeft(5)
	root.Left.InsertRight(3)
	root.Right.InsertRight(9)
	fmt.Println(largestValInEachTreeRow(root)) // [1, 3, 9]
}
