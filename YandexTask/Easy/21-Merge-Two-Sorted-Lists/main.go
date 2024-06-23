package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	var res, tmp *ListNode
	if list1.Val <= list2.Val {
		tmp = list1
		res = tmp
		list1 = list1.Next
	} else {
		tmp = list2
		res = tmp
		list2 = list2.Next
	}
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tmp.Next = list2
			break
		} else if list2 == nil {
			tmp.Next = list1
			break
		} else {
			if list1.Val <= list2.Val {
				tmp.Next = list1
				list1 = list1.Next
			} else {
				tmp.Next = list2
				list2 = list2.Next
			}
		}
		tmp = tmp.Next
	}

	return res
}
