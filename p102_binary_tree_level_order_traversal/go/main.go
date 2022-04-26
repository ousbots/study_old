package binary_tree_level_order_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	results := [][]int{}
	parents := []*TreeNode{root}

	for len(parents) != 0 {
		values := []int{}
		length := len(parents)

		for i := 0; i < length; i++ {
			if parents[0] != nil {
				values = append(values, parents[0].Val)
				parents = append(parents, parents[0].Left)
				parents = append(parents, parents[0].Right)
			}
			parents = parents[1:]
		}

		if len(values) > 0 {
			results = append(results, values)
		}
	}

	return results
}
