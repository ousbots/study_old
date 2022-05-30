package p62_unique_paths

import "testing"

func TestProvided(t *testing.T) {
	valid := 28
	if ans := uniquePaths(3, 7); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 3
	if ans := uniquePaths(3, 2); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	valid := 1
	if ans := uniquePaths(1, 1); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 1
	if ans := uniquePaths(1, 2); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 1
	if ans := uniquePaths(2, 1); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 2
	if ans := uniquePaths(2, 2); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 3
	if ans := uniquePaths(2, 3); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}

	valid = 3
	if ans := uniquePaths(3, 2); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}

func TestExtra(t *testing.T) {
	valid := 10518300
	if ans := uniquePaths(25, 9); ans != valid {
		t.Fatalf("%v is not %v", ans, valid)
	}
}
