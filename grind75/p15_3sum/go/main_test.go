package _3sum

import (
	"sort"
	"testing"
)

func TestProvided(t *testing.T) {
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	if ans := threeSum([]int{-1, 0, 1, 2, -1, -4}); !comp(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	expected = [][]int{}
	if ans := threeSum([]int{}); !comp(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	expected = [][]int{}
	if ans := threeSum([]int{0}); !comp(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}
}

func TestBasic(t *testing.T) {
	expected := [][]int{{-1, -1, 2}}
	if ans := threeSum([]int{-1, 1, -1, 1, -1, 1, -1, 2}); !comp(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	expected = [][]int{{1, 1, -2}}
	if ans := threeSum([]int{1, 1, 1, -2, -2, 1, 1, -2, 1, -2}); !comp(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for outer := range a {
		if len(a[outer]) != len(b[outer]) {
			return false
		}

		sort.Ints(a[outer])
		sort.Ints(b[outer])

		for inner := range a[outer] {
			if a[outer][inner] != b[outer][inner] {
				return false
			}
		}
	}

	return true
}
