package number_of_islands

type Pair struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '1' {
				count += 1
				sink(grid, x, y)
			}
		}
	}

	return count
}

func sink(grid [][]byte, x, y int) {
	visit := []Pair{{x, y}}

	for len(visit) != 0 {
		check := visit[len(visit)-1]
		visit = visit[:len(visit)-1]

		if grid[check.y][check.x] == '1' {
			grid[check.y][check.x] = '0'

			if check.x > 0 {
				visit = append(visit, Pair{x: check.x - 1, y: check.y})
			}

			if check.x < len(grid[check.y])-1 {
				visit = append(visit, Pair{x: check.x + 1, y: check.y})
			}

			if check.y > 0 {
				visit = append(visit, Pair{x: check.x, y: check.y - 1})
			}

			if check.y < len(grid)-1 {
				visit = append(visit, Pair{x: check.x, y: check.y + 1})
			}
		}
	}
}
