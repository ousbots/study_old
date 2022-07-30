package p79_word_search

import "testing"

func TestBasic(t *testing.T) {
	board := [][]byte{{'a'}}
	word := "a"
	valid := true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{{'a'}}
	word = "b"
	valid = false
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{{'a', 'b', 'c', 'd', 'e'}}
	word = "bcd"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{{'a'}, {'b'}, {'c'}, {'d'}, {'e'}}
	word = "bcd"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	word = "abehi"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	word = "abehgd"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	word = "abcfihgde"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	word = "edghifcba"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}
}

func TestExtra(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCESEEEFS"
	valid := true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCEFSADEESE"
	valid = true
	if ans := exist(board, word); ans != valid {
		t.Fatalf("expected %t for %s", valid, word)
	}
}
