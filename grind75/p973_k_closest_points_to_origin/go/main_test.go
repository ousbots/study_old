package k_closest_points_to_origin

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	valid := [][]int{{-2, 2}}
	if ans := kClosest([][]int{{1, 3}, {-2, 2}}, 1); !compare(ans, valid) {
		t.Fatalf("expected %v, not %v", valid, ans)
	}

	valid = [][]int{{3, 3}, {-2, 4}}
	if ans := kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2); !compare(ans, valid) {
		t.Fatalf("expected %v, not %v", valid, ans)
	}
}

func TestMore(t *testing.T) {
	valid := [][]int{{-2, 2}, {2, -2}}
	if ans := kClosest([][]int{{1, 3}, {-2, 2}, {2, -2}}, 2); !compare(ans, valid) {
		t.Fatalf("expected %v, not %v", valid, ans)
	}

	valid = [][]int{{1, 1}}
	if ans := kClosest([][]int{{2,2},{2,2},{2,2},{2,2},{2,2},{2,2},{1,1}}, 1); !compare(ans, valid) {
		t.Fatalf("expected %v, not %v", valid, ans)
	}
}

func compare(a, b [][]int) bool {
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] == a[j][0] {
			return a[i][1] < a[j][1]
		}

		return a[i][0] < a[j][0]
	})

	sort.Slice(b, func(i, j int) bool {
		if b[i][0] == b[j][0] {
			return b[i][1] < b[j][1]
		}

		return b[i][0] < b[j][0]
	})

	if len(a) != len(b) {
		return false
	}

	for outer := range a {
		if len(a[outer]) != len(b[outer]) {
			return false
		}

		for inner := range a[outer] {
			if a[outer][inner] != b[outer][inner] {
				return false
			}
		}
	}

	return true
}
