package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	tmp := head
	firstTmp, secondTmp := &ListNode{}, &ListNode{}
	first, second := firstTmp, secondTmp
	for tmp != nil {
		if tmp.Val < x {
			firstTmp.Next = tmp
			tmp = tmp.Next
			firstTmp = firstTmp.Next
		} else {
			secondTmp.Next = tmp
			tmp = tmp.Next
			secondTmp = secondTmp.Next
		}
	}
	secondTmp.Next = nil
	firstTmp.Next = second.Next
	return first.Next
}
