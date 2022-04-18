package flood_fill

import "testing"

func TestProvided(t *testing.T) {
	expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	if ans := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

	expected = [][]int{{2, 2, 2}, {2, 2, 2}}
	if ans := floodFill([][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}
}

func TestBasic(t *testing.T) {
	expected := [][]int{{1, 0, 1}, {0, 2, 0}, {1, 0, 1}}
	if ans := floodFill([][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}}, 1, 1, 2); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

	expected = [][]int{{0, 2, 0}, {2, 2, 2}, {0, 2, 0}}
	if ans := floodFill([][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 1, 1, 2); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

	expected = [][]int{{2, 2, 0}, {0, 2, 2}, {2, 2, 2}}
	if ans := floodFill([][]int{{1, 1, 0}, {0, 1, 1}, {1, 1, 1}}, 0, 0, 2); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}

	expected = [][]int{{0, 0, 0}, {0, 1, 1}}
	if ans := floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1); !check(ans, expected) {
		t.Fatalf("expected %v not %v", expected, ans)
	}
}

func check(a, b [][]int) bool {
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
