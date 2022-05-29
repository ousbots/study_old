package p54_spiral_matrix

import "testing"

func TestBasic(t *testing.T) {
	input := [][]int{{1}}
	valid := []int{1}
	if ans := spiralOrder(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	valid = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if ans := spiralOrder(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{
		{1, 2, 3},
		{6, 5, 4},
	}
	valid = []int{1, 2, 3, 4, 5, 6}
	if ans := spiralOrder(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{{1}, {2}, {3}}
	valid = []int{1, 2, 3}
	if ans := spiralOrder(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b []int) bool {
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