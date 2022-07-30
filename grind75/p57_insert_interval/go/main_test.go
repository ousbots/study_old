package insert_interval

import "testing"

func TestProvided(t *testing.T) {
	ans := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	expected := [][]int{{1, 5}, {6, 9}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	expected = [][]int{{1, 2}, {3, 10}, {12, 16}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}
}

func TestBasic(t *testing.T) {
	ans := insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{7, 7})
	expected := [][]int{{0, 2}, {4, 6}, {7, 7}, {8, 10}, {12, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{11, 11})
	expected = [][]int{{0, 2}, {4, 6}, {8, 10}, {11, 11}, {12, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{7, 9})
	expected = [][]int{{0, 2}, {4, 6}, {7, 10}, {12, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{9, 11})
	expected = [][]int{{0, 2}, {4, 6}, {8, 11}, {12, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{5, 13})
	expected = [][]int{{0, 2}, {4, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{-1, 7})
	expected = [][]int{{-1, 7}, {8, 10}, {12, 14}, {16, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{11, 19})
	expected = [][]int{{0, 2}, {4, 6}, {8, 10}, {11, 19}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{0, 2}, {4, 6}, {8, 10}, {12, 14}, {16, 18}}, []int{1, 17})
	expected = [][]int{{0, 18}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{1, 5}}, []int{6, 8})
	expected = [][]int{{1, 5}, {6, 8}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
	}

	ans = insert([][]int{{6, 8}}, []int{1, 5})
	expected = [][]int{{1, 5}, {6, 8}}
	if !compare(ans, expected) {
		t.Fatalf("expected %v, not %v", expected, ans)
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
