package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var leight int
	tmp := head
	for tmp != nil {
		leight++
		tmp = tmp.Next
	}
	lenf := leight - k%leight
	if lenf == leight {
		return head
	}
	tmp = head
	for i := 0; i < lenf-1; i++ {
		tmp = tmp.Next
	}
	rev := tmp.Next
	tmp.Next = nil
	result := rev
	for i := 0; i < leight-lenf-1; i++ {
		rev = rev.Next
	}
	rev.Next = head
	return result
}
