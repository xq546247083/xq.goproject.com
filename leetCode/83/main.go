package main

import "fmt"

func main() {
	fmt.Println(deleteDuplicates(nil))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	temp := head
	for {
		if temp == nil {
			break
		}

		if temp.Next == nil {
			break
		}

		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
	}

	return head
}
