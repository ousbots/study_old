package rotting_oranges

type Pair struct {
	x int
	y int
}

func orangesRottingExtraMem(grid [][]int) int {
	time := 0
	changed := []Pair{}
	seen := make(map[Pair]bool)

	count := 0
	for {
		count++
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == 2 {
					changed = append(changed, rotExtraMem(grid, seen, y, x)...)
				}
			}
		}

		if len(changed) != 0 {
			for _, pair := range changed {
				grid[pair.y][pair.x] = 2
			}

			changed = changed[0:0]
			time += 1
		} else {
			break
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				return -1
			}
		}
	}

	return time
}

func rotExtraMem(grid [][]int, seen map[Pair]bool, y, x int) []Pair {
	change := []Pair{}
	visit := []Pair{{x, y}}

	for len(visit) != 0 {
		check := visit[len(visit)-1]
		visit = visit[:len(visit)-1]

		if !seen[check] && grid[check.y][check.x] == 2 {
			seen[check] = true

			if check.y > 0 {
				visit = append(visit, Pair{x: check.x, y: check.y - 1})
			}

			if check.y < len(grid)-1 {
				visit = append(visit, Pair{x: check.x, y: check.y + 1})
			}

			if check.x > 0 {
				visit = append(visit, Pair{x: check.x - 1, y: check.y})
			}

			if check.x < len(grid[check.y])-1 {
				visit = append(visit, Pair{x: check.x + 1, y: check.y})
			}
		}

		if grid[check.y][check.x] == 1 {
			change = append(change, Pair{x: check.x, y: check.y})
		}
	}

	return change
}

func orangesRotting(grid [][]int) int {
	time := 0
	changed := []Pair{}

	count := 0
	for {
		count++
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == 2 {
					changed = append(changed, rot(grid, y, x)...)
				}
			}
		}

		if len(changed) != 0 {
			for _, pair := range changed {
				grid[pair.y][pair.x] = 2
			}

			changed = changed[0:0]
			time += 1
		} else {
			break
		}
	}

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 1 {
				return -1
			}
		}
	}

	return time
}

func rot(grid [][]int, y, x int) []Pair {
	change := []Pair{}
	visit := []Pair{{x, y}}

	for len(visit) != 0 {
		check := visit[len(visit)-1]
		visit = visit[:len(visit)-1]

		if grid[check.y][check.x] == 2 {
			grid[check.y][check.x] = 3

			if check.y > 0 {
				visit = append(visit, Pair{x: check.x, y: check.y - 1})
			}

			if check.y < len(grid)-1 {
				visit = append(visit, Pair{x: check.x, y: check.y + 1})
			}

			if check.x > 0 {
				visit = append(visit, Pair{x: check.x - 1, y: check.y})
			}

			if check.x < len(grid[check.y])-1 {
				visit = append(visit, Pair{x: check.x + 1, y: check.y})
			}
		}

		if grid[check.y][check.x] == 1 {
			change = append(change, Pair{x: check.x, y: check.y})
		}
	}

	return change
}
