package lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	small := p.Val
	big := q.Val
	if small > big {
		small = q.Val
		big = p.Val
	}

	current := root
	for current != nil {
		if current.Val >= small && current.Val <= big {
			break
		}

		if current.Val == small || current.Val == big {
			break
		}

		if current.Val > big {
			current = current.Left
		} else if current.Val < small {
			current = current.Right
		}
	}

	return current
}
