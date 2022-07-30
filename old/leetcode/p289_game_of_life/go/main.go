package game_of_life

func gameOfLife(board [][]int) {
	orig := make([][]int, len(board))
	for i, row := range board {
		orig[i] = append(orig[i], row...)
	}
	
	for y, row := range board {
		for x, val := range row {
			neighbors := neighbors(orig, x, y)
			
			if val == 0 {
				if neighbors == 3 {
					board[y][x] = 1
				}
			} else {
				if neighbors < 2 || neighbors > 3 {
					board[y][x] = 0
				}
			}
		}
	}
}

func neighbors(board [][]int, x, y int) int {
	if len(board) == 0 {
		return 0
	}

	count := 0

	if x > 0 {
		if board[y][x-1] == 1 {
			count++
		}
		
		if y > 0 {
			if board[y-1][x-1] == 1 {
				count++
			}
		}
	
		if y < len(board) - 1 {
			if board[y+1][x-1] == 1 {
				count++
			}
		}
	}
	
		
	if x < len(board[0]) - 1 {
		if board[y][x+1] == 1 {
			count++
		}
			
		if y > 0 {
			if board[y-1][x+1] == 1 {
				count++
			}
		}
		
		if y < len(board) - 1 {
			if board[y+1][x+1] == 1 {
				count++
			}
		}
	}

	if y > 0 {
		if board[y-1][x] == 1 {
			count++
		}
	}

		
	if y < len(board) - 1 {
		if board[y+1][x] == 1 {
			count++
		}
	}

	return count
}
