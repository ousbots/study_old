package p236_lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestorSlow(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := contains(root.Left, p) || contains(root.Left, q)
	right := contains(root.Right, p) || contains(root.Right, q)
	self := root.Val == p.Val || root.Val == q.Val

	if (left && right) || (self && left) || (self && right) {
		return root
	} else if left {
		return lowestCommonAncestorSlow(root.Left, p, q)
	} else {
		return lowestCommonAncestorSlow(root.Right, p, q)
	}
}

func contains(root, element *TreeNode) bool {
	if root == nil || element == nil {
		return false
	}

	if root.Val == element.Val {
		return true
	}

	return contains(root.Left, element) || contains(root.Right, element)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left == nil && right == nil {
		return nil
	}

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}
}
