package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func lenghtList(head *ListNode) int {
	lenght := 0
	tmp := head
	for tmp != nil {
		lenght++
		tmp = tmp.Next
	}
	return lenght
}

func reverseList(head *ListNode) *ListNode {
	var res, next_val *ListNode
	tmp := head
	for tmp != nil {
		next_val = tmp.Next
		tmp.Next = res
		res = tmp
		tmp = next_val
	}
	return res
}

func isPalindrome(head *ListNode) bool {
	lenght := lenghtList(head)
	tmp := head
	for i := 0; i < (lenght+1)/2; i++ {
		tmp = tmp.Next
	}
	first := head
	second := reverseList(tmp)
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first, second = first.Next, second.Next
	}
	return true
}
