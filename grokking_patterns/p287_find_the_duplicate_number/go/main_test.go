package find_the_duplicate_number

import "testing"

type testCase struct {
	nums     []int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []testCase{
		{nums: []int{1, 1}, expected: 1},
		{nums: []int{1, 3, 4, 2, 2}, expected: 2},
		{nums: []int{3, 1, 3, 4, 2}, expected: 3},
	}

	for _, test := range tests {
		ans := findDuplicate(test.nums)
		if ans != test.expected {
			t.Fatalf("case %v expected %d not %d", test.nums, test.expected, ans)
		}
	}
}
