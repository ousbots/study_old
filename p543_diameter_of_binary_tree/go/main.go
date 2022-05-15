package diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	_ = dfs(root, &max)

	return max
}

func dfs(root *TreeNode, max *int) int {
	if root == nil {
		return -1
	}

	left, right := dfs(root.Left, max)+1, dfs(root.Right, max)+1

	if sum := left + right; sum > *max {
		*max = sum
	}

	if left > right {
		return left
	} else {
		return right
	}
}
