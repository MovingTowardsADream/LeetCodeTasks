package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && symmetric(left.Right, right.Left) && symmetric(left.Left, right.Right)
}

// or
//
//var result bool
//
//func isSymmetric(root *TreeNode) bool {
//	result = true
//	if root == nil {
//		return true
//	}
//	symmetric(root.Left, root.Right)
//	return result
//}
//
//func symmetric(left, right *TreeNode) {
//	if left == nil && right == nil {
//		return
//	}
//	if left == nil || right == nil || left.Val != right.Val {
//		result = false
//		return
//	}
//	symmetric(left.Right, right.Left)
//	symmetric(left.Left, right.Right)
//}
//
// or
//
//func isSymmetric(root *TreeNode) bool {
//	if root == nil { return true }
//	res := true
//
//	var permSymmetric func(*TreeNode, *TreeNode)
//	permSymmetric = func(left, right *TreeNode) {
//		if left == nil && right == nil { return }
//
//		if !res || left == nil || right == nil || left.Val != right.Val {
//			res = false
//			return
//		}
//
//		permSymmetric(left.Right, right.Left)
//		permSymmetric(left.Left, right.Right)
//	}
//
//	permSymmetric(root.Left, root.Right)
//
//	return res
//}
