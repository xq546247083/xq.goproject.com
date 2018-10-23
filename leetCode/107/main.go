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
	a4.Val = 4
	a5.Val = 5

	a1.Left = a2
	a1.Right = a4
	a2.Right = a3
	a4.Right = a5

	fmt.Println(levelOrderBottom(a1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0, 32)
	result = levelOrderBottomV2(result, 0, root)

	lenResult := len(result)
	finalResult := make([][]int, lenResult, lenResult)
	// 反转顺序
	for index, item := range result {
		finalResult[lenResult-index-1] = item
	}

	return finalResult
}

// 递归添加子元素
func levelOrderBottomV2(list [][]int, level int, node *TreeNode) [][]int {
	// 如果当前节点为空，直接返回
	if node == nil {
		return list
	}

	if len(list) <= level {
		list = append(list, []int{})
	}

	// 把当前节点添加到当前等级的数组后面
	list[level] = append(list[level], node.Val)

	// 追加下一等级的左右节点
	if node.Left != nil {
		list = levelOrderBottomV2(list, level+1, node.Left)
	}
	if node.Right != nil {
		list = levelOrderBottomV2(list, level+1, node.Right)
	}

	return list
}
