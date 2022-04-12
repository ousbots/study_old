package game_of_life

import "testing"

func TestProvided(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}}
	ans := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}}
	gameOfLife(board)
	
	if !comp(board, ans) {
		t.Fatalf("failed: expected %v found %v", ans, board)
	}

	board = [][]int{{1, 1}, {1, 0}}
	ans = [][]int{{1, 1}, {1, 1}}
	gameOfLife(board)
	
	if !comp(board, ans) {
		t.Fatalf("failed: expected %v found %v", ans, board)
	}
}

func TestBasic(t *testing.T) {
	board := [][]int{{0}}
	ans := [][]int{{0}}
	gameOfLife(board)
	
	if !comp(board, ans) {
		t.Fatalf("failed: expected %v found %v", ans, board)
	}

	board = [][]int{{1}}
	ans = [][]int{{0}}
	gameOfLife(board)
	
	if !comp(board, ans) {
		t.Fatalf("failed: expected %v found %v", ans, board)
	}
}


func comp(first, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}

	for i, firstRow := range first {
		secondRow := second[i]

		if len(firstRow) != len(secondRow) {
			return false
		}
			
		for index := range firstRow {
			if firstRow[index] != secondRow[index] {
				return false
			}
		}
	}
	
	return true
}
