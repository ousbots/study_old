package _3sum_closest

import "testing"

type TestCase struct {
	nums   []int
	target int
	output int
}

func TestBasic(t *testing.T) {
	tests := []TestCase{
		{nums: []int{0, 0, 1}, target: 1, output: 1},
		{nums: []int{1, 0, 0, 1}, target: 1, output: 1},
		{nums: []int{-2, 2, -1, 5, -1, -2}, target: 1, output: 1},
		{nums: []int{5, -1, 2, -1}, target: 1, output: 0},
		{nums: []int{-1, 2, 1, -4}, target: 1, output: 2},
		{nums: []int{0, 0, 0}, target: 1, output: 0},
	}

	for _, test := range tests {
		ans := threeSumClosest(test.nums, test.target)
		if ans != test.output {
			t.Fatalf("for %v with target %d expected %d not %d", test.nums, test.target, test.output, ans)
		}
	}
}
