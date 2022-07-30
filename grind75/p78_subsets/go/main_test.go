package p78_subsets

import (
	"sort"
	"testing"
)

func TestBasic(t *testing.T) {
	input := []int{1}
	valid := [][]int{{}, {1}}
	if ans := subsets(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = []int{1, 2}
	valid = [][]int{{}, {1}, {2}, {1, 2}}
	if ans := subsets(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = []int{1, 2, 3, 4, 5}
	valid = [][]int{
		{},
		{1}, {2}, {3}, {4}, {5},
		{1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 3}, {2, 4}, {2, 5}, {3, 4}, {3, 5}, {4, 5},
		{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
		{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	if ans := subsets(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

// Inefficient way of comparing vs. filtering one of the arrays.
func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		sort.Ints(a[i])
		sort.Ints(b[i])
	}

	sort.Slice(a, func(i, j int) bool {
		if len(a[i]) == len(a[j]) {
			for k := range a[i] {
				if a[i][k] != a[j][k] {
					return a[i][k] < a[j][k]
				}
			}
			return false
		} else {
			return len(a[i]) < len(a[j])
		}
	})

	sort.Slice(b, func(i, j int) bool {
		if len(b[i]) == len(b[j]) {
			for k := range b[i] {
				if b[i][k] != b[j][k] {
					return b[i][k] < b[j][k]
				}
			}
			return false
		} else {
			return len(b[i]) < len(b[j])
		}
	})

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
