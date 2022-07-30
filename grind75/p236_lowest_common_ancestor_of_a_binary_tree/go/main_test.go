package p236_lowest_common_ancestor_of_a_binary_tree

import "testing"

func TestProvided(t *testing.T) {
	input := &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}}
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 5}, &TreeNode{Val: 1}); ans == nil || ans.Val != 3 {
		t.Fatalf("%d is not 3", ans.Val)
	}

	if ans := lowestCommonAncestor(input, &TreeNode{Val: 5}, &TreeNode{Val: 4}); ans == nil || ans.Val != 5 {
		t.Fatalf("%d is not 5", ans.Val)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	if ans := lowestCommonAncestor(input, &TreeNode{Val: 1}, &TreeNode{Val: 2}); ans == nil || ans.Val != 1 {
		t.Fatalf("%d is not 1", ans.Val)
	}
}
