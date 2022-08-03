package sort_colors

import "testing"

type TestCase struct {
	nums     []int
	expected []int
}

func checkSlices(a, b []int) bool {
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

func TestBasic(t *testing.T) {
	tests := []TestCase{
		{nums: []int{0}, expected: []int{0}},
		{nums: []int{2, 1, 0}, expected: []int{0, 1, 2}},
		{nums: []int{2, 0, 1}, expected: []int{0, 1, 2}},
		{nums: []int{1, 2, 0}, expected: []int{0, 1, 2}},
		{nums: []int{2, 0, 2, 1, 1, 0}, expected: []int{0, 0, 1, 1, 2, 2}},
		{nums: []int{2, 1, 0, 1, 0, 0, 2, 2, 1}, expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
	}

	for _, test := range tests {
		sortColors(test.nums)
		if !checkSlices(test.nums, test.expected) {
			t.Fatalf("%v and %v do not match", test.nums, test.expected)
		}
	}
}
