package binary_search

import "testing"

func TestProvided(t *testing.T) {
	if ans := search([]int{-1, 0, 3, 5, 9, 12}, 9); ans != 4 {
		t.Fatalf("expected 4 not %d", ans)
	}

	if ans := search([]int{-1, 0, 3, 5, 9, 12}, 2); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	if ans := search([]int{1}, 1); ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}

	if ans := search([]int{1}, 2); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}

	if ans := search([]int{-1, 1}, 1); ans != 1 {
		t.Fatalf("expected 1 not %d", ans)
	}

	if ans := search([]int{-1, 1}, -1); ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}

	if ans := search([]int{-1, 1}, 2); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}
}
