package balanced_binary_tree

import "testing"

func TestProvided(t *testing.T) {
	input := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if isBalanced(input) != true {
		t.Fatal("expected true")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}
	if isBalanced(input) == true {
		t.Fatal("expected false")
	}

	if isBalanced(nil) != true {
		t.Fatal("expected true")
	}
}

func TestBasic(t *testing.T) {
	input := &TreeNode{Val: 1}
	if isBalanced(input) != true {
		t.Fatal("expected true")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	if isBalanced(input) == true {
		t.Fatal("expected false")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}}}
	if isBalanced(input) == true {
		t.Fatal("expected true")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	if isBalanced(input) != true {
		t.Fatal("expected true")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	if isBalanced(input) == true {
		t.Fatal("expected false")
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}}
	if isBalanced(input) == true {
		t.Fatal("expected false")
	}

}
