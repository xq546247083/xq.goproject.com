package main

import "fmt"

func main() {
	a1 := new(TreeNode)
	a2 := new(TreeNode)
	a3 := new(TreeNode)
	a4 := new(TreeNode)
	a5 := new(TreeNode)
	a1.Val = 1
	a2.Val = 2
	a3.Val = 3
	a4.Val = 2
	a5.Val = 3

	a1.Left = a2
	a2.Right = a3
	a1.Right = a4
	a4.Right = a5

	fmt.Println(maxDepth(a1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}
