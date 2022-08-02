package max_consecutive_ones_iii

import "testing"

type TestCase struct {
	nums   []int
	k      int
	output int
}

func TestBasic(t *testing.T) {
	tests := []TestCase{
		{nums: []int{0}, k: 0, output: 0},
		{nums: []int{0}, k: 1, output: 1},
		{nums: []int{0, 1, 0}, k: 0, output: 1},
		{nums: []int{0, 1, 0}, k: 1, output: 2},
		{nums: []int{0, 0, 1, 1, 1, 0, 0}, k: 0, output: 3},
		{nums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, k: 2, output: 6},
		{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, k: 3, output: 10},
		{nums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, k: 2, output: 11},
	}

	for _, test := range tests {
		ans := longestOnes(test.nums, test.k)
		if ans != test.output {
			t.Fatalf("expected %d for %v not %d", test.output, test.nums, ans)
		}
	}
}
