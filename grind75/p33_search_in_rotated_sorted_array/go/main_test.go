package search_in_rotated_sorted_array

import "testing"

func TestProvided(t *testing.T) {
	if ans := search([]int{4, 5, 6, 7, 0, 1, 2}, 0); ans != 4 {
		t.Fatalf("%d is not %d", ans, 4)
	}

	if ans := search([]int{4, 5, 6, 7, 0, 1, 2}, 3); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{1}, 0); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}
}

func TestBasic(t *testing.T) {
	if ans := search([]int{0, 1, 2, 3, 4}, 2); ans != 2 {
		t.Fatalf("%d is not %d", ans, 2)
	}

	if ans := search([]int{0, 1, 2, 3, 4}, 0); ans != 0 {
		t.Fatalf("%d is not %d", ans, 0)
	}

	if ans := search([]int{0, 1, 2, 3, 4}, 4); ans != 4 {
		t.Fatalf("%d is not %d", ans, 4)
	}

	if ans := search([]int{0, 1, 2, 3, 4}, -1); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{0, 1, 2, 3, 4}, 5); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{1, 2, 3, 4, 0}, 2); ans != 1 {
		t.Fatalf("%d is not %d", ans, 1)
	}

	if ans := search([]int{1, 2, 3, 4, 0}, 0); ans != 4 {
		t.Fatalf("%d is not %d", ans, 4)
	}

	if ans := search([]int{1, 2, 3, 4, 0}, 4); ans != 3 {
		t.Fatalf("%d is not %d", ans, 3)
	}

	if ans := search([]int{1, 2, 3, 4, 0}, -1); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{1, 2, 3, 4, 0}, 5); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, 4); ans != 0 {
		t.Fatalf("%d is not %d", ans, 0)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, 3); ans != 4 {
		t.Fatalf("%d is not %d", ans, 4)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, 0); ans != 1 {
		t.Fatalf("%d is not %d", ans, 1)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, 2); ans != 3 {
		t.Fatalf("%d is not %d", ans, 3)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, -1); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}

	if ans := search([]int{4, 0, 1, 2, 3}, 5); ans != -1 {
		t.Fatalf("%d is not %d", ans, -1)
	}
}