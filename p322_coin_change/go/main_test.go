package coin_change

import "testing"

func TestProvided(t *testing.T) {
	if ans := coinChange([]int{1, 2, 5}, 11); ans != 3 {
		t.Fatalf("expected 3, not %d", ans)
	}

	if ans := coinChange([]int{2}, 3); ans != -1 {
		t.Fatalf("expected -1, not %d", ans)
	}
	if ans := coinChange([]int{1}, 0); ans != 0 {
		t.Fatalf("expected 0, not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	if ans := coinChange([]int{5, 1, 2}, 11); ans != 3 {
		t.Fatalf("expected 3, not %d", ans)
	}

	if ans := coinChange([]int{1}, 1); ans != 1 {
		t.Fatalf("expected 0, not %d", ans)
	}

	if ans := coinChange([]int{1, 3, 4}, 6); ans != 2 {
		t.Fatalf("expected 2, not %d", ans)
	}
}
