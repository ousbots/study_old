package _4sum

import "testing"

type TestCase struct {
	nums     []int
	target   int
	expected [][]int
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			for j := range a[i] {
				if a[i][j] != b[i][j] {
					return false
				}
			}
		}
	}

	return true
}

func TestBasic(t *testing.T) {
	tests := []TestCase{
		{nums: []int{1, 0, -1, 0, -2, 2}, target: 0, expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{nums: []int{2, 2, 2, 2, 2}, target: 8, expected: [][]int{{2, 2, 2, 2}}},
	}

	for _, test := range tests {
		ans := fourSum(test.nums, test.target)
		if !comp(ans, test.expected) {
			t.Fatalf("nums %v target %d should be %v not %v", test.nums, test.target, test.expected, ans)
		}
	}
}
