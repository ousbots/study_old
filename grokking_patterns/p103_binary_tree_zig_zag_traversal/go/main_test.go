package binary_tree_zig_zag_level_order_traversal

import "testing"

type testCase struct {
	head     *TreeNode
	expected [][]int
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		{head: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: [][]int{{3}, {20, 9}, {15, 7}}},
		{head: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}}, expected: [][]int{{1}, {3, 2}, {4, 5}}},
		{head: &TreeNode{Val: 1}, expected: [][]int{{1}}},
		{head: nil, expected: [][]int{}},
	}

	for _, test := range tests {
		ans := zigzagLevelOrder(test.head)
		if !comp(ans, test.expected) {
			t.Fatalf("expected %v not %v", test.expected, ans)
		}
	}
}