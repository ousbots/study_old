package rotting_oranges

import "testing"

func TestProvided(t *testing.T) {
	input := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	valid := 4
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	valid = -1
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{0, 2}}
	valid = 0
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	input := [][]int{{2}}
	valid := 0
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{1}}
	valid = -1
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{0}}
	valid = 0
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{1, 2}}
	valid = 1
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 1}}
	valid = 4
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 2}}
	valid = 2
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]int{{2, 0, 2}, {0, 1, 0}, {2, 0, 2}}
	valid = -1
	if ans := orangesRotting(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}
}