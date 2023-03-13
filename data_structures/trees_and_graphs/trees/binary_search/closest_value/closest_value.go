package main

import (
	"fmt"
	"math"
	tn "trees_and_graphs"
)

func closestValue(root *tn.TreeNode, target float64) int {
	val, closest := root.Val, root.Val

	for root != nil {
		val = root.Val
		if math.Abs(float64(val)-target) < math.Abs(float64(closest)-target) {
			closest = val
		}

		if int(target) < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closest
}

func main() {
	root := tn.New(4)
	root.InsertLeft(2)
	root.InsertRight(5)
	root.Left.InsertLeft(1)
	root.Left.InsertRight(3)
	fmt.Println(closestValue(root, 3.7142)) // 4
}
