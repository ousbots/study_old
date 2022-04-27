package evaluate_reverse_polish_notation

import "testing"

func TestProvided(t *testing.T) {
	if ans := evalRPN([]string{"2", "1", "+", "3", "*"}); ans != 9 {
		t.Fatalf("expected 9, not %d", ans)
	}

	if ans := evalRPN([]string{"4", "13", "5", "/", "+"}); ans != 6 {
		t.Fatalf("expected 6, not %d", ans)
	}

	if ans := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}); ans != 22 {
		t.Fatalf("expected 22, not %d", ans)
	}
}
