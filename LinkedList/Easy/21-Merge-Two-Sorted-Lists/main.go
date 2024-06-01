package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp.Next = l2
			tmp, l2 = tmp.Next, l2.Next
		} else if l2 == nil || l1.Val < l2.Val {
			tmp.Next = l1
			tmp, l1 = tmp.Next, l1.Next
		} else {
			tmp.Next = l2
			tmp, l2 = tmp.Next, l2.Next
		}
	}
	return result.Next
}
