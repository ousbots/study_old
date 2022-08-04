package subarray_product_less_than_k

import "testing"

type TestCase struct {
	nums     []int
	k        int
	expected int
}

func TestBasic(t *testing.T) {
	tests := []TestCase{
		{nums: []int{10, 5, 2, 6}, k: 100, expected: 8},
		{nums: []int{1, 2, 3}, k: 0, expected: 0},
		{nums: []int{0, 2, 3, 4}, k: 1, expected: 4},
	}

	for _, test := range tests {
		ans := numSubarrayProductLessThanK(test.nums, test.k)
		if ans != test.expected {
			t.Fatalf("testing nums=%v k=%d expected %d not %d", test.nums, test.k, test.expected, ans)
		}
	}
}
