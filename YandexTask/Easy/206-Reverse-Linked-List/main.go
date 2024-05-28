package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var res, val *ListNode
	tmp := head
	for tmp != nil {
		val = tmp.Next
		tmp.Next = res
		res = tmp
		tmp = val
	}
	return res
}
