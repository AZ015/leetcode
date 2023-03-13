package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func dfs(node *tn.TreeNode, maxSoFar int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, int(math.Max(float64(maxSoFar), float64(node.Val))))
	right := dfs(node.Right, int(math.Max(float64(maxSoFar), float64(node.Val))))

	ans := left + right
	if node.Val > maxSoFar {
		ans += 1
	}

	return ans
}

func goodNodes(root *tn.TreeNode) int {
	return dfs(root, math.MinInt)
}

func main() {
	root := tn.New(5)
	root.InsertLeft(4)
	root.InsertRight(8)
	root.Left.InsertLeft(11)
	root.Left.Left.InsertLeft(7)
	root.Left.Left.InsertRight(2)
	root.Right.InsertRight(4)
	root.Right.Right.InsertRight(1)
	root.Right.InsertLeft(13)

	fmt.Println(goodNodes(root))
}
