package clone_graph

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	seen := make(map[int]*Node)
	return clone(node, seen)
}

// Depth first search (clone neighbors as needed).
func clone(node *Node, seen map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := seen[node.Val]; !ok {
		newNode := Node{Val: node.Val}
		seen[node.Val] = &newNode

		for _, elem := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, clone(elem, seen))
		}
	}

	return seen[node.Val]
}
