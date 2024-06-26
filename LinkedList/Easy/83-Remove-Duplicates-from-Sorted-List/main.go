package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	left, right := head, head
	for right != nil {
		if left.Val == right.Val {
			right = right.Next
		} else {
			left.Next = right
			left = right
		}
	}
	left.Next = nil
	return head
}
