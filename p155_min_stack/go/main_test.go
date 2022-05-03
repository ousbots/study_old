package min_stack

import "testing"

func TestProvided(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if ans := stack.GetMin(); ans != -3 {
		t.Fatalf("%d is not -3", ans)
	}

	stack.Pop()

	if ans := stack.Top(); ans != 0 {
		t.Fatalf("%d is not 0", ans)
	}

	if ans := stack.GetMin(); ans != -2 {
		t.Fatalf("%d is not -2", ans)
	}
}

func TestBasic(t *testing.T) {
	stack := Constructor()
	stack.Push(-1)
	stack.Push(0)
	stack.Push(1)

	if ans := stack.GetMin(); ans != -1 {
		t.Fatalf("%d is not -1", ans)
	}
	if ans := stack.Top(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}

	stack.Push(-5)

	if ans := stack.GetMin(); ans != -5 {
		t.Fatalf("%d is not -5", ans)
	}
	if ans := stack.Top(); ans != -5 {
		t.Fatalf("%d is not -5", ans)
	}

	stack.Pop()

	if ans := stack.GetMin(); ans != -1 {
		t.Fatalf("%d is not -1", ans)
	}
	if ans := stack.Top(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}

	stack.Pop()

	if ans := stack.GetMin(); ans != -1 {
		t.Fatalf("%d is not -1", ans)
	}
	if ans := stack.Top(); ans != 0 {
		t.Fatalf("%d is not 0", ans)
	}

	stack.Pop()

	if ans := stack.GetMin(); ans != -1 {
		t.Fatalf("%d is not -1", ans)
	}
	if ans := stack.Top(); ans != -1 {
		t.Fatalf("%d is not -1", ans)
	}

	stack = Constructor()
	stack.Push(1)
	stack.Push(2)

	if ans := stack.GetMin(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}
	if ans := stack.Top(); ans != 2 {
		t.Fatalf("%d is not 2", ans)
	}

	stack.Push(-5)

	if ans := stack.GetMin(); ans != -5 {
		t.Fatalf("%d is not -5", ans)
	}
	if ans := stack.Top(); ans != -5 {
		t.Fatalf("%d is not -5", ans)
	}

	stack.Pop()

	if ans := stack.GetMin(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}
	if ans := stack.Top(); ans != 2 {
		t.Fatalf("%d is not 2", ans)
	}

	stack.Pop()

	if ans := stack.GetMin(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}
	if ans := stack.Top(); ans != 1 {
		t.Fatalf("%d is not 1", ans)
	}

	stack.Pop()

	stack.Push(5)
	if ans := stack.GetMin(); ans != 5 {
		t.Fatalf("%d is not 5", ans)
	}
}