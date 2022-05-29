package p199_binary_tree_right_side_view

import "testing"

func TestBasic(t *testing.T) {
	valid := []int{}
	if ans := rightSideView(nil); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input := &TreeNode{Val: 1}
	valid = []int{1}
	if ans := rightSideView(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	valid = []int{1, 2}
	if ans := rightSideView(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	valid = []int{1, 3}
	if ans := rightSideView(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	valid = []int{1, 2, 3}
	if ans := rightSideView(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}}
	valid = []int{1, 3, 4, 5}
	if ans := rightSideView(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}