package path_sum_ii

import "testing"

type testCase struct {
	head      *TreeNode
	targetSum int
	expected  [][]int
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
		{head: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}}}, targetSum: 22, expected: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{head: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, targetSum: 5, expected: [][]int{}},
		{head: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, targetSum: 0, expected: [][]int{}},
		{head: &TreeNode{Val: 1, Left: &TreeNode{Val: -1}}, targetSum: 0, expected: [][]int{{1, -1}}},
		{head: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}, targetSum: 3, expected: [][]int{{1, 2}, {1, 2}}},
	}

	for _, test := range tests {
		ans := pathSum(test.head, test.targetSum)
		if !comp(ans, test.expected) {
			t.Fatalf("expected %v not %v for sum %d", test.expected, ans, test.targetSum)
		}
	}
}
