package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result, first, second := head, head, head
	for second != nil {
		if first.Val == second.Val {
			second = second.Next
		} else {
			first.Next = second
			first = first.Next
		}
	}
	first.Next = second
	return result
}
