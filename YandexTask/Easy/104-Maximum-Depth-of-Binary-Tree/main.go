package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maximum := 0
	var main func(node *TreeNode, sum int)
	main = func(node *TreeNode, sum int) {
		if node == nil {
			maximum = max(maximum, sum)
			return
		}
		main(node.Left, sum+1)
		main(node.Right, sum+1)
	}
	main(root, 0)
	return maximum
}
