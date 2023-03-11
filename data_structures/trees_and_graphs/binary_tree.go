package trees_and_graphs

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func (tn *TreeNode) InsertLeft(val int) {
	tn.Left = New(val)
}

func (tn *TreeNode) InsertRight(val int) {
	tn.Right = New(val)
}

func PreorderDfs(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	PreorderDfs(node.Left)
	PreorderDfs(node.Right)
}

func InorderDfs(node *TreeNode) {
	if node == nil {
		return
	}

	PreorderDfs(node.Left)
	fmt.Println(node.Val)
	PreorderDfs(node.Right)
}

func PostorderDfs(node *TreeNode) {
	if node == nil {
		return
	}

	PreorderDfs(node.Left)
	PreorderDfs(node.Right)
	fmt.Println(node.Val)
}

func main() {
	root := New(0)
	root.InsertLeft(1)
	root.InsertRight(2)
	root.Left.InsertLeft(3)
	root.Left.InsertRight(4)
	root.Right.InsertRight(5)
	root.Left.Left.InsertLeft(6)
	root.Left.Right.InsertLeft(7)
	root.Left.Right.InsertRight(8)

	PreorderDfs(root)
	fmt.Println("----------------------------")
	InorderDfs(root)
	fmt.Println("----------------------------")
	PostorderDfs(root)
}
