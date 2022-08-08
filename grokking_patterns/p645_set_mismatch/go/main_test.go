package set_mismatch

import (
	"sort"
	"testing"
)

type testCase struct {
	nums     []int
	expected []int
}

func comp(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		{nums: []int{1, 2, 2, 4}, expected: []int{2, 3}},
		{nums: []int{1, 1}, expected: []int{1, 2}},
	}

	for _, test := range tests {
		ans := findErrorNums(test.nums)
		if !comp(ans, test.expected) {
			t.Fatalf("case %v expected %v not %v", test.nums, test.expected, ans)
		}
	}
}
