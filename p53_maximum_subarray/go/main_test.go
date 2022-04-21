package maximum_subarray

import "testing"

func TestProvided(t *testing.T) {
	if ans := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); ans != 6 {
		t.Fatalf("expected 6 not %d", ans)
	}

	if ans := maxSubArray([]int{5, 4, -1, 7, 8}); ans != 23 {
		t.Fatalf("expected 23 not %d", ans)
	}

	if ans := maxSubArray([]int{1}); ans != 1 {
		t.Fatalf("expected 1 not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	if ans := maxSubArray([]int{-1}); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}

	if ans := maxSubArray([]int{0, -1}); ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}

	if ans := maxSubArray([]int{-1, 0}); ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}

	if ans := maxSubArray([]int{-1, 0, -1}); ans != 0 {
		t.Fatalf("expected 0 not %d", ans)
	}

	if ans := maxSubArray([]int{-1, -2, -3}); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}

	if ans := maxSubArray([]int{-2, -1, -3}); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}

	if ans := maxSubArray([]int{-3, -2, -1}); ans != -1 {
		t.Fatalf("expected -1 not %d", ans)
	}
}
