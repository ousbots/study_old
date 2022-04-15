package invert_binary_tree

import (
	"strconv"
	"testing"
)

func TestProvided(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil}, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6, Left: nil, Right: nil}, Right: &TreeNode{Val: 9, Left: nil, Right: nil}}}
	root = invertTree(root)
	ans := &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 6, Left: nil, Right: nil}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: &TreeNode{Val: 1, Left: nil, Right: nil}}}

	if !compare(root, ans) {
		t.Fatalf("%v incorrect, expected %v", root, ans)
	}

}

func TestBasic(t *testing.T) {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}
	root = invertTree(root)
	ans := &TreeNode{Val: 4, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}
	if !compare(root, ans) {
		t.Fatalf("%v incorrect, expected %v", root, ans)
	}

	root = &TreeNode{Val: 4, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: nil}}
	root = invertTree(root)
	ans = &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: nil, Right: nil}, Right: nil}
	if !compare(root, ans) {
		t.Fatalf("%v incorrect, expected %v", root, ans)
	}

	root = nil
	root = invertTree(root)
	if root != nil {
		t.Fatalf("%v incorrect, expected nil", root)
	}

	root = &TreeNode{Val: 1, Left: nil, Right: nil}
	root = invertTree(root)
	ans = &TreeNode{Val: 1, Left: nil, Right: nil}
	if !compare(root, ans) {
		t.Fatalf("%v incorrect, expected %v", root, ans)
	}
}

func compare(first, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}

	if first.Val != second.Val {
		return false
	}

	return compare(first.Left, second.Left) && compare(first.Right, second.Right)
}

func toString(root *TreeNode) string {
	var helper func(*TreeNode) string
	helper = func(node *TreeNode) string {
		if node == nil {
			return ""
		}

		return " " + strconv.Itoa(node.Val) + helper(node.Left) + helper(node.Right)
	}

	return "[" + helper(root) + "]"
}