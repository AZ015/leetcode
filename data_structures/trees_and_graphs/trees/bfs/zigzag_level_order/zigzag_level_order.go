package main

import (
	"fmt"
	"sort"
	tn "trees_and_graphs"
)

func zigzagLevelOrder(root *tn.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	q := tn.NewQueue[*tn.TreeNode]()
	q.Push(root)
	reverse := false

	for !q.IsEmpty() {
		levelArr := make([]int, 0)
		for i := q.Length() - 1; i >= 0; i-- {
			node := q.PopLeft()
			levelArr = append(levelArr, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}

			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		if reverse {
			sort.Sort(sort.Reverse(sort.IntSlice(levelArr)))
		}
		reverse = !reverse
		ans = append(ans, levelArr)
	}

	return ans
}

func main() {
	root := tn.New(3)
	root.InsertLeft(9)
	root.InsertRight(20)
	root.Right.InsertLeft(15)
	root.Right.InsertRight(7)
	fmt.Println(zigzagLevelOrder(root)) // [[3], [20,9],[15,7]]
}
