package diameter_of_binary_tree

import "testing"

func TestProvided(t *testing.T) {
	if ans := diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}); ans != 3 {
		t.Fatalf("%d is not %d", ans, 3)
	}

	if ans := diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}); ans != 1 {
		t.Fatalf("%d is not %d", ans, 1)
	}
}

func TestBasic(t *testing.T) {
	if ans := diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}}, Right: &TreeNode{Val: 3}}); ans != 4 {
		t.Fatalf("%d is not %d", ans, 4)
	}
}
