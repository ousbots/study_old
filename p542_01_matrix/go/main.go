package _01_matrix

type Point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {
	dist := make([][]int, len(mat))
	queue := []Point{}

	for y := range mat {
		dist[y] = make([]int, len(mat[y]))
		for x := range mat[y] {
			if mat[y][x] == 0 {
				queue = append(queue, Point{x: x, y: y})
			} else {
				dist[y][x] = -1
			}
		}
	}

	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]

		value := dist[tmp.y][tmp.x] + 1

		if tmp.x > 0 {
			if dist[tmp.y][tmp.x-1] == -1 {
				dist[tmp.y][tmp.x-1] = value
				queue = append(queue, Point{x: tmp.x - 1, y: tmp.y})
			} else {
				if value < dist[tmp.y][tmp.x-1] {
					dist[tmp.y][tmp.x-1] = value
				}
			}
		}

		if tmp.x < len(mat[0])-1 {
			if dist[tmp.y][tmp.x+1] == -1 {
				dist[tmp.y][tmp.x+1] = value
				queue = append(queue, Point{x: tmp.x + 1, y: tmp.y})
			} else {
				if value < dist[tmp.y][tmp.x+1] {
					dist[tmp.y][tmp.x+1] = value
				}
			}
		}

		if tmp.y > 0 {
			if dist[tmp.y-1][tmp.x] == -1 {
				dist[tmp.y-1][tmp.x] = value
				queue = append(queue, Point{x: tmp.x, y: tmp.y - 1})
			} else {
				if value < dist[tmp.y-1][tmp.x] {
					dist[tmp.y-1][tmp.x] = value
				}
			}
		}

		if tmp.y < len(mat)-1 {
			if dist[tmp.y+1][tmp.x] == -1 {
				dist[tmp.y+1][tmp.x] = value
				queue = append(queue, Point{x: tmp.x, y: tmp.y + 1})
			} else {
				if value < dist[tmp.y+1][tmp.x] {
					dist[tmp.y+1][tmp.x] = value
				}
			}
		}
	}

	return dist
}
