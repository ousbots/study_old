package binary_tree_level_order_traversal_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeTuple struct {
	Node  *TreeNode
	Level int
}

// Returns a slice where every element is a slice of nodes found at that level in the tree. As in,
// the values in the slice at position i of the parent slice, represent the values of all nodes at
// a depth of i in the tree ordered from left to right.
//
// This solution uses a BFS to parse all the nodes on the same level in order. The queue used for
// the BFS stores both the node and the level the node was found at. That enables knowing which
// position to put the node's value into the ordered slice.
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	order := [][]int{}

	queue := []NodeTuple{{root, 0}}
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Node.Left != nil {
			queue = append(queue, NodeTuple{current.Node.Left, current.Level + 1})
		}

		if current.Node.Right != nil {
			queue = append(queue, NodeTuple{current.Node.Right, current.Level + 1})
		}

		if len(order) > current.Level {
			order[current.Level] = append(order[current.Level], current.Node.Val)
		} else {
			order = append(order, []int{current.Node.Val})
		}
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	return order
}

// Returns a slice of slices where each element slice is values at a given level ordered from left
// to right in the tree. The slice of slices is ordered from the bottom up of the tree.
//
// This solution uses a DFS helper function that goes down left before right and fills in the
// levels map of the left to right order of the values on each level of the tree. The map is then
// converted to a slice of slices where each slice is one of the map values. The slice is then
// reversed and returned.
func levelOrderBottomDFS(root *TreeNode) [][]int {
	levels := make(map[int][]int)
	dive(root, 0, levels)

	order := make([][]int, len(levels))

	for level, vals := range levels {
		order[level] = vals
	}

	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}

	return order
}

// Depth first search of the tree left to right that builds the levels map of level -> values.
func dive(root *TreeNode, level int, levels map[int][]int) {
	if root == nil {
		return
	}

	dive(root.Left, level+1, levels)
	dive(root.Right, level+1, levels)

	levels[level] = append(levels[level], root.Val)
}
