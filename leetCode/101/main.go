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

	fmt.Println(isSymmetric(a1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftNodeList := make([]*TreeNode, 0, 32)
	rightNodeList := make([]*TreeNode, 0, 32)

	leftNodeList = addNodeByLeft(leftNodeList, root.Left)
	rightNodeList = addNodeByRight(rightNodeList, root.Right)

	if len(leftNodeList) != len(rightNodeList) {
		return false
	}

	// 判断是否完全相同，相同则对称
	for index, item := range leftNodeList {
		if rightNodeList[index] == nil && item == nil {
			continue
		}

		if rightNodeList[index] != nil && item != nil && rightNodeList[index].Val == item.Val {
			continue
		}

		return false
	}

	return true
}

// 左节点，一直从左节点递归添加节点
func addNodeByLeft(list []*TreeNode, node *TreeNode) []*TreeNode {
	list = append(list, node)
	if node != nil {
		// 递归添加
		list = addNodeByLeft(list, node.Left)
		list = addNodeByLeft(list, node.Right)
	}

	return list
}

// 右节点，一直从右节点递归添加节点
func addNodeByRight(list []*TreeNode, node *TreeNode) []*TreeNode {
	list = append(list, node)
	if node != nil {
		// 递归添加
		list = addNodeByRight(list, node.Right)
		list = addNodeByRight(list, node.Left)
	}

	return list
}
