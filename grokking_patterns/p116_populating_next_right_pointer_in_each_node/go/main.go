package populating_next_right_pointers_in_each_node

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Bundle struct {
	Node  *Node
	Level int
}

// Returns the root of the tree after each node on the same depth level is connected from left to
// right (via the Node.Next field).
//
// This solution uses a DFS to make sure that all nodes of the same depth level will be processed
// sequentially. To connect the nodes, the last node checked is kept so that if the last and
// current nodes in the DFS queue have the same level they can be connected so that the last points
// to the current.
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []Bundle{{Node: root, Level: 0}}
	last := (*Bundle)(nil)

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Node.Left != nil {
			queue = append(queue, Bundle{Node: current.Node.Left, Level: current.Level + 1})
		}

		if current.Node.Right != nil {
			queue = append(queue, Bundle{Node: current.Node.Right, Level: current.Level + 1})
		}

		if last != nil {
			if last.Level == current.Level {
				last.Node.Next = current.Node
			}
		}

		last = &current
	}

	return root
}
