package balanced_binary_tree

// Binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return distance(root) != -1
}

func distance(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := distance(root.Left)
	if left == -1 {
		return -1
	}

	right := distance(root.Right)
	if right == -1 {
		return right
	}

	total := left - right
	if total > 1 || total < -1 {
		return -1
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
