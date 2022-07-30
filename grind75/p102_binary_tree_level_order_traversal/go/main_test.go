package binary_tree_level_order_traversal

import "testing"

func TestProvided(t *testing.T) {
	input := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	if ans := levelOrder(input); !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	input = &TreeNode{Val: 1}
	expected = [][]int{{1}}
	if ans := levelOrder(input); !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	input = nil
	expected = [][]int{}
	if ans := levelOrder(input); !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}
}

func TestBasic(t *testing.T) {
	input := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}}
	expected := [][]int{{1}, {2, 4}, {3, 5, 6}}
	if ans := levelOrder(input); !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}}}}}
	expected = [][]int{{1}, {2}, {3}, {4}, {5}}
	if ans := levelOrder(input); !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}
}

func compare(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for outer := range a {
		if len(a[outer]) != len(b[outer]) {
			return false
		}

		for inner := range a[outer] {
			if a[outer][inner] != b[outer][inner] {
				return false
			}
		}
	}

	return true
}
