package main

import "fmt"

func main() {
	a1 := new(TreeNode)
	a2 := new(TreeNode)
	a3 := new(TreeNode)
	a1.Val = 1
	a2.Val = 2
	a3.Val = 3
	a1.Left = a2
	a1.Right = a3

	b1 := new(TreeNode)
	b2 := new(TreeNode)
	b3 := new(TreeNode)
	b1.Val = 1
	b2.Val = 2
	b3.Val = 4
	b1.Left = b2
	b1.Right = b3

	fmt.Println(isSameTree(a1, b1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	} else {
		if !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
			return false
		} else {
			return true
		}
	}
}
