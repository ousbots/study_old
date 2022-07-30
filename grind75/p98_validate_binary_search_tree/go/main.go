package validate_binary_search_tree

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root.Left, math.MinInt32-1, root.Val) && check(root.Right, root.Val, math.MaxInt32+1)
}

func check(node *TreeNode, lo, hi int) bool {
	if node == nil {
		return true
	}

	return lo < node.Val && node.Val < hi && check(node.Left, lo, node.Val) && check(node.Right, node.Val, hi)
}
