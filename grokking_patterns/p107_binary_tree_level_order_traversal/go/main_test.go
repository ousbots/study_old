package binary_tree_level_order_traversal_ii

import "testing"

type testCase struct {
	root     *TreeNode
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
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, expected: [][]int{{15, 7}, {9, 20}, {3}}},
		{root: &TreeNode{Val: 1}, expected: [][]int{{1}}},
		{root: nil, expected: [][]int{}},
	}

	for _, test := range tests {
		ans := levelOrderBottom(test.root)
		if !comp(ans, test.expected) {
			t.Fatalf("expected %v not %v", test.expected, ans)
		}
	}
}
