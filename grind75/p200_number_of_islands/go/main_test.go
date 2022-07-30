package number_of_islands

import "testing"

func TestProvided(t *testing.T) {
	input := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	valid := 1
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	valid = 3
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}
}

func TestBasic(t *testing.T) {
	input := [][]byte{{'1', '1', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '0', '1', '0', '1'}, {'1', '0', '1', '1', '1'}}
	valid := 1
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]byte{{'1', '0', '1', '0', '1'}, {'0', '1', '0', '1', '0'}, {'1', '0', '1', '0', '1'}, {'0', '1', '0', '1', '0'}}
	valid = 10
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]byte{{'1', '1', '0', '0', '1', '1'}, {'1', '1', '0', '0', '1', '1'}, {'0', '0', '1', '1', '0', '0'}, {'0', '0', '1', '1', '0', '0'}, {'1', '1', '0', '0', '1', '1'}, {'1', '1', '0', '0', '1', '1'}}
	valid = 5
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]byte{{'1'}}
	valid = 1
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}

	input = [][]byte{{'0'}}
	valid = 0
	if ans := numIslands(input); ans != valid {
		t.Fatalf("%d is not %d", ans, valid)
	}
}
