package p104_maximum_depth_of_binary_tree

import "testing"

func TestProvided(t *testing.T) {
	input := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if ans := maxDepth(input); ans != 3 {
		t.Fatalf("%d is not 3", ans)
	}

	input = &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	if ans := maxDepth(input); ans != 2 {
		t.Fatalf("%d is not 2", ans)
	}
}

func TestBasic(t *testing.T) {
	input := &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Right: &TreeNode{Val: 99, Left: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if ans := maxDepth(input); ans != 4 {
		t.Fatalf("%d is not 4", ans)
	}
}
