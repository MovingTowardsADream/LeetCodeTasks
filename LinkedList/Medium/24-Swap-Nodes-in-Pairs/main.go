package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	list := head
	first := &ListNode{}
	result := first
	var second *ListNode
	for list != nil && list.Next != nil {
		second = list.Next.Next
		tmp := list.Next
		tmp.Next = list
		first.Next = tmp
		tmp.Next.Next = second
		first = tmp.Next
		list = second
	}
	return result.Next
}

func main() {

}
