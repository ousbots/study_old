package first_bad_version

import "testing"

func TestProvided(t *testing.T) {
	bad_version = 4
	if ans := firstBadVersion(5); ans != 4 {
		t.Fatalf("expected 4, not %d", ans)
	}

	bad_version = 1
	if ans := firstBadVersion(1); ans != 1 {
		t.Fatalf("expected 1, not %d", ans)
	}
}

func TestBasic(t *testing.T) {
	bad_version = 1
	if ans := firstBadVersion(99); ans != 1 {
		t.Fatalf("expected 1, not %d", ans)
	}

	bad_version = 999
	if ans := firstBadVersion(999); ans != 999 {
		t.Fatalf("expected 999, not %d", ans)
	}

	bad_version = 432
	if ans := firstBadVersion(999); ans != 432 {
		t.Fatalf("expected 432, not %d", ans)
	}
}
