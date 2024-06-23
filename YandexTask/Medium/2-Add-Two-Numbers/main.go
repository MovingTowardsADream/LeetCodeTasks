package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, tmp *ListNode
	remains := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			remains += l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			remains += l1.Val
			l1 = l1.Next
		} else {
			remains += l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		}
		if tmp == nil {
			tmp = &ListNode{remains % 10, nil}
			res = tmp
		} else {
			tmp.Next = &ListNode{remains % 10, nil}
			tmp = tmp.Next
		}
		remains /= 10
	}
	if remains > 0 {
		tmp.Next = &ListNode{remains, nil}
	}
	return res
}
