package p199_binary_tree_right_side_view

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	right := append([]int{root.Val}, rightSideView(root.Right)...)
	left := rightSideView(root.Left)

	if len(left) >= len(right) {
		right = append(right, left[len(right)-1:]...)
	}

	return right
}
