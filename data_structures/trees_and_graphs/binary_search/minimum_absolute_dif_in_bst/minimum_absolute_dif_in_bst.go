package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func dfs(node *tn.TreeNode) []int {
	if node == nil {
		return []int{}
	}

	left := dfs(node.Left)
	right := dfs(node.Right)

	left = append(left, node.Val)
	left = append(left, right...)

	return left
}

func minimumAbsoluteDifInBst(root *tn.TreeNode) int {
	values := dfs(root)
	ans := math.MaxInt
	for i := 1; i < len(values); i++ {
		ans = int(math.Min(float64(ans), float64(values[i]-values[i-1])))
	}

	return ans
}

func main() {
	root := tn.New(9)
	root.InsertLeft(5)
	root.InsertRight(15)
	root.Left.InsertLeft(1)
	root.Left.InsertRight(7)
	fmt.Println(minimumAbsoluteDifInBst(root)) //
}
