package implement_queue_using_stacks

import "testing"

func TestProvided(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)

	if ans := queue.Peek(); ans != 1 {
		t.Fatalf("expected 1, not %d\n", ans)
	}

	if ans := queue.Pop(); ans != 1 {
		t.Fatalf("expected 1, not %d\n", ans)
	}

	if ans := queue.Empty(); ans != false {
		t.Fatalf("expected false, not %v\n", ans)
	}
}

func TestBasic(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if ans := queue.Peek(); ans != 1 {
		t.Fatalf("expected 3, not %d\n", ans)
	}

	if ans := queue.Pop(); ans != 1 {
		t.Fatalf("expected 1, not %d\n", ans)
	}

	if ans := queue.Peek(); ans != 2 {
		t.Fatalf("expected 2, not %d\n", ans)
	}

	queue.Push(4)

	if ans := queue.Peek(); ans != 2 {
		t.Fatalf("expected 2, not %d\n", ans)
	}

	queue.Push(5)

	if ans := queue.Peek(); ans != 2 {
		t.Fatalf("expected 2, not %d\n", ans)
	}

	if ans := queue.Empty(); ans != false {
		t.Fatalf("expected false, not %v", ans)
	}

	if ans := queue.Pop(); ans != 2 {
		t.Fatalf("expected 2, not %d\n", ans)
	}

	if ans := queue.Peek(); ans != 3 {
		t.Fatalf("expected 3, not %d\n", ans)
	}

	if ans := queue.Pop(); ans != 3 {
		t.Fatalf("expected 3, not %d\n", ans)
	}

	if ans := queue.Peek(); ans != 4 {
		t.Fatalf("expected 4, not %d\n", ans)
	}

	if ans := queue.Pop(); ans != 4 {
		t.Fatalf("expected 4, not %d\n", ans)
	}

	if ans := queue.Peek(); ans != 5 {
		t.Fatalf("expected 5, not %d\n", ans)
	}

	if ans := queue.Empty(); ans != false {
		t.Fatalf("expected false, not %v", ans)
	}

	if ans := queue.Pop(); ans != 5 {
		t.Fatalf("expected 5, not %d\n", ans)
	}

	if ans := queue.Empty(); ans != true {
		t.Fatalf("expected true, not %v", ans)
	}
}
