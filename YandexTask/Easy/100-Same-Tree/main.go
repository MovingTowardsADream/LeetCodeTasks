package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(first *TreeNode, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}
	if first == nil || second == nil || first.Val != second.Val {
		return false
	}
	return isSameTree(first.Left, second.Left) && isSameTree(first.Right, second.Right)
}
