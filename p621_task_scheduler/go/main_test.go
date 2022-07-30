package task_scheduler

import "testing"

func TestProvided(t *testing.T) {
	ans := leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2)
	expected := 8
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0)
	expected = 6
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2)
	expected = 16
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}
}

func TestBasic(t *testing.T) {
	ans := leastInterval([]byte{'A'}, 0)
	expected := 1
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A'}, 9)
	expected = 1
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A'}, 0)
	expected = 2
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A'}, 1)
	expected = 3
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A'}, 2)
	expected = 4
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'B'}, 2)
	expected = 2
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A', 'B'}, 0)
	expected = 3
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A', 'B'}, 1)
	expected = 3
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}

	ans = leastInterval([]byte{'A', 'A', 'B'}, 2)
	expected = 4
	if ans != expected {
		t.Fatalf("expected %d not %d", expected, ans)
	}
}
