package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return validationTree(root.Left, MinInt, root.Val, true) && validationTree(root.Right, root.Val, MaxInt, true)
}

func validationTree(node *TreeNode, minV, maxV int, result bool) bool {
	if node == nil {
		return true
	}
	if node.Val > minV && node.Val < maxV {
		return validationTree(node.Left, minV, min(node.Val, maxV), result) && validationTree(node.Right, max(node.Val, minV), maxV, result) && result
	}
	return false
}
