package p79_word_search

func exist(board [][]byte, word string) bool {
	for y := range board {
		for x := range board[0] {
			if board[y][x] == word[0] {
				if find(board, word, make(map[int]bool), y, x, 0) {
					return true
				}
			}
		}
	}

	return false
}

func find(board [][]byte, word string, seen map[int]bool, y, x, index int) bool {
	key := (100 * y) + x
	if seen[key] || index >= len(word) || board[y][x] != word[index] {
		return false
	}

	if index == len(word)-1 {
		return true
	}

	seen[key] = true

	if x < len(board[0])-1 {
		if find(board, word, seen, y, x+1, index+1) {
			return true
		}
	}

	if x > 0 {
		if find(board, word, seen, y, x-1, index+1) {
			return true
		}
	}

	if y < len(board)-1 {
		if find(board, word, seen, y+1, x, index+1) {
			return true
		}
	}

	if y > 0 {
		if find(board, word, seen, y-1, x, index+1) {
			return true
		}
	}

	delete(seen, key)

	return false
}

func existBad(board [][]byte, word string) bool {
	for y := range board {
		for x := range board[y] {
			if board[y][x] == word[0] {
				if findBad(board, word, make(map[int]bool), y, x, 0) {
					return true
				}
			}
		}
	}

	return false
}

func findBad(board [][]byte, word string, seen map[int]bool, y, x, index int) bool {
	if index >= len(word) || y >= len(board) || y < 0 || x >= len(board[0]) || x < 0 || board[y][x] != word[index] {
		return false
	}

	// m, n <= 6 so y*100 + x is unique for all values of x and y.
	key := (y * 100) + x
	if seen[key] {
		return false
	}
	seen[key] = true

	if index == len(word)-1 {
		return true
	}

	return findBad(board, word, copy(seen), y, x+1, index+1) ||
		findBad(board, word, copy(seen), y, x-1, index+1) ||
		findBad(board, word, copy(seen), y+1, x, index+1) ||
		findBad(board, word, copy(seen), y-1, x, index+1)
}

// Oof, this is bad, but the board dimensions are small enough that it should be ok.
func copy(a map[int]bool) map[int]bool {
	b := make(map[int]bool, len(a))
	for key, val := range a {
		b[key] = val
	}

	return b
}
