package climbing_stairs

import "testing"

func TestProvided(t *testing.T) {
	expected := 2
	if ans := climbStairs(2); ans != expected {
		t.Fatalf("%d is not %d", ans, expected)
	}

	expected = 3
	if ans := climbStairs(3); ans != expected {
		t.Fatalf("%d is not %d", ans, expected)
	}
}
