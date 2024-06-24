package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func lenghtList(head *ListNode) int {
	tmp := head
	lenght := 0
	for tmp != nil {
		lenght++
		tmp = tmp.Next
	}
	return lenght
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	lenght := lenghtList(head)
	if lenght == n {
		return head.Next
	}
	tmp := head
	for i := 1; i < lenght-n; i++ {
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return head
}
