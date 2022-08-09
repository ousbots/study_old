package binary_tree_zig_zag_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeLevel struct {
	Node  *TreeNode
	Level int
}

// Returns a slice where each element of the slice is a list of values of all nodes at that level
// in the tree in zig-zag order (alternating left-right right-left). As in, the slice at position i
// in the returned slice has the values of all nodes from level i in the tree in the correct order.
//
// This solution uses a BFS to go through every node in the same level in order. The depth of the
// level is also stored in the queue to know which position in the order slice to put the value in.
// In order to create the zig-zag order,
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	order := [][]int{}
	queue := []NodeLevel{{Node: root, Level: 0}}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Node.Left != nil {
			queue = append(queue, NodeLevel{Node: current.Node.Left, Level: current.Level + 1})
		}

		if current.Node.Right != nil {
			queue = append(queue, NodeLevel{Node: current.Node.Right, Level: current.Level + 1})
		}

		if len(order) <= current.Level {
			order = append(order, []int{current.Node.Val})
		} else {
			order[current.Level] = append(order[current.Level], current.Node.Val)
		}
	}

	for level := range order {
		if level%2 == 1 {
			for i, j := 0, len(order[level])-1; i < j; i, j = i+1, j-1 {
				order[level][i], order[level][j] = order[level][j], order[level][i]
			}
		}
	}

	return order
}
