package p105_construct_binary_tree_from_preorder_and_inorder_traversal

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "null"
	}

	return strconv.Itoa(t.Val) + " " + t.Left.String() + " " + t.Right.String()
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := TreeNode{Val: preorder[0]}

	if len(preorder) == 0 {
		return &root
	}

	mid := 0
	for i := range inorder {
		if inorder[i] == root.Val {
			mid = i
			break
		}
	}

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return &root
}
