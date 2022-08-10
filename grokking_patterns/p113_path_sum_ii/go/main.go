package path_sum_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Returns a slice of root to leaf paths from the given tree that add up to the targetSum.
//
// This solution uses a DFS to find the root to leaf paths that add up to the targeted sum. The
// helper function is used to recurse down while maintaining the current sum and path. The base
// condition is being a leaf node (no children) and for those nodes if the current sum equals the
// target, then the path is returned and appended with other results as the recursion unwinds.
func pathSum(root *TreeNode, targetSum int) [][]int {
	return helper(root, targetSum, 0, []int{})
}

// Helper function to perform DFS search.
func helper(root *TreeNode, targetSum int, sum int, path []int) [][]int {
	if root == nil {
		return nil
	}

	sum += root.Val
	path = append(path, root.Val)

	if root.Left == nil && root.Right == nil {
		if sum == targetSum {
			return [][]int{path}
		} else {
			return nil
		}
	}

	temp := make([]int, len(path))
	copy(temp, path)

	left := helper(root.Left, targetSum, sum, path)
	right := helper(root.Right, targetSum, sum, temp)

	result := [][]int{}
	result = append(result, left...)
	result = append(result, right...)

	return result
}
