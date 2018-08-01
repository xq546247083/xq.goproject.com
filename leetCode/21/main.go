package main

import (
	"fmt"
)

func main() {
	l11 := &ListNode{Val: 1}
	// l12 := &ListNode{Val: 3}
	// l14 := &ListNode{Val: 4}
	// l11.Next = l12
	// l12.Next = l14

	l21 := &ListNode{Val: 3}
	// l23 := &ListNode{Val: 2}
	// l24 := &ListNode{Val: 4}
	// l21.Next = l23
	// l23.Next = l24

	result := mergeTwoLists(l11, l21)
	for {
		fmt.Println(result)

		if result.Next == nil {
			break
		}

		result = result.Next
	}
}

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 由于是大小有序的链表
// 所以把大链表往小的链表合并
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		tempNode := l1
		l1 = l2
		l2 = tempNode
	}

	result := l2
	for {
		// 插入失败，且l2到尾了，那么，接上就可以了
		if l2.Next == nil {
			l2.Next = l1
			return result
		}

		// 如果插入成功，则插入下一个
		flag, nextNode := ha(l1, l2)
		if flag {
			// 如果没有下一个，说明全部插入了
			if nextNode == nil {
				return result
			}
			l1 = nextNode
			continue
		}

		l2 = l2.Next
	}
}

func ha(l1 *ListNode, l2 *ListNode) (bool, *ListNode) {
	tempNode := l2.Next
	result := l1.Next

	if l1.Val >= l2.Val && l1.Val <= tempNode.Val {
		l2.Next = l1
		l1.Next = tempNode
		return true, result
	}

	return false, nil
}
