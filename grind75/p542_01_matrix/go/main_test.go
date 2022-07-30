package _01_matrix

import "testing"

func TestProvided(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	expected := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	if ans := updateMatrix(input); !compare(ans, expected) {
		t.Fatalf("input %v -> expected %v, not %v", input, expected, ans)
	}

	input = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	expected = [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}
	if ans := updateMatrix(input); !compare(ans, expected) {
		t.Fatalf("input %v -> expected %v, not %v", input, expected, ans)
	}
}

func TestBasic(t *testing.T) {
	input := [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}
	expected := [][]int{{4, 3, 2}, {3, 2, 1}, {2, 1, 0}}
	if ans := updateMatrix(input); !compare(ans, expected) {
		t.Fatalf("input %v -> expected %v, not %v", input, expected, ans)
	}

	input = [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}
	expected = [][]int{{0, 1, 2}, {1, 2, 1}, {2, 1, 0}}
	if ans := updateMatrix(input); !compare(ans, expected) {
		t.Fatalf("input %v -> expected %v, not %v", input, expected, ans)
	}

	input = [][]int{{1, 1, 0}, {1, 1, 1}, {1, 1, 0}}
	expected = [][]int{{2, 1, 0}, {3, 2, 1}, {2, 1, 0}}
	if ans := updateMatrix(input); !compare(ans, expected) {
		t.Fatalf("input %v -> expected %v, not %v", input, expected, ans)
	}
}

func compare(a, b [][]int) bool {
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
