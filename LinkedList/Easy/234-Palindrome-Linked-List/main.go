package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
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

func middle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	first := head
	second := reverse(middle(head))
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first, second = first.Next, second.Next
	}
	return true
}
