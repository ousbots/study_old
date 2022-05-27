package p56_merge_intervals

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	valid := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if ans := merge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	input := [][]int{{1, 3}, {5, 6}, {10, 14}, {20, 30}}
	valid := [][]int{{1, 3}, {5, 6}, {10, 14}, {20, 30}}
	if ans := merge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{{1, 3}, {5, 6}, {10, 14}, {0, 13}}
	valid = [][]int{{0, 14}}
	if ans := merge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{{1, 3}, {5, 6}, {10, 14}, {4, 9}}
	valid = [][]int{{1, 3}, {4, 9}, {10, 14}}
	if ans := merge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}

	input = [][]int{{1, 3}, {5, 6}, {10, 14}, {1, 10}}
	valid = [][]int{{1, 14}}
	if ans := merge(input); !comp(ans, valid) {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == b[i][0] {
			return a[i][1] < a[j][1]
		} else {
			return a[i][0] < b[i][0]
		}
	})

	sort.Slice(b, func(i, j int) bool {
		if b[i][0] == b[i][0] {
			return b[i][1] < b[j][1]
		} else {
			return b[i][0] < b[i][0]
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
