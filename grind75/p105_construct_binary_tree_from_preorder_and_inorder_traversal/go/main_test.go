package p105_construct_binary_tree_from_preorder_and_inorder_traversal

import "testing"

func TestProvided(t *testing.T) {
	valid := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if ans := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	valid := &TreeNode{Val: 3}
	if ans := buildTree([]int{3}, []int{3}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = &TreeNode{Val: 3, Left: &TreeNode{Val: 20}}
	if ans := buildTree([]int{3, 20}, []int{20, 3}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = &TreeNode{Val: 3, Right: &TreeNode{Val: 20}}
	if ans := buildTree([]int{3, 20}, []int{3, 20}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = &TreeNode{Val: 3, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if ans := buildTree([]int{3, 20, 15, 7}, []int{3, 15, 20, 7}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = &TreeNode{Val: 3, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if ans := buildTree([]int{3, 20, 15, 7}, []int{15, 20, 7, 3}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestExtra(t *testing.T) {
	valid := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	if ans := buildTree([]int{1, 2, 3}, []int{3, 2, 1}); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && comp(a.Left, b.Left) && comp(a.Right, b.Right)
}
