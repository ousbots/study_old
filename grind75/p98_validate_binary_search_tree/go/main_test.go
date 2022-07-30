package validate_binary_search_tree

import "testing"

func TestProvided(t *testing.T) {
	input := TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	if isValidBST(&input) != true {
		t.Fatal("expected true")
	}

	input = TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	if isValidBST(&input) != false {
		t.Fatal("expected false")
	}
}

func TestBasic(t *testing.T) {
	input := TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if isValidBST(&input) != false {
		t.Fatal("expected false")
	}

	input = TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	if isValidBST(&input) != false {
		t.Fatal("expected false")
	}

	input = TreeNode{Val: 5, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 7}}}
	if isValidBST(&input) != false {
		t.Fatal("expected false")
	}

	input = TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}
	if isValidBST(&input) != false {
		t.Fatal("expected false")
	}

	input = TreeNode{Val: -2147483648, Right: &TreeNode{Val: 2147483647}}
	if isValidBST(&input) != true {
		t.Fatal("expected true")
	}

}
