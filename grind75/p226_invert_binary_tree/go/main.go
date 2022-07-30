package invert_binary_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	swapNodes(root)

	return root
}

func swapNodes(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	swapNodes(root.Left)
	swapNodes(root.Right)
}

func invertTreeNoRecurse(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]

		node.Right, node.Left = node.Left, node.Right

		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}

		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
	}

	return root
}
