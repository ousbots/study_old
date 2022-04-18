package flood_fill

type position struct {
	x int
	y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]

	if color == newColor {
		return image
	}

	queue := []position{{x: sc, y: sr}}

	for len(queue) > 0 {
		// A stack is faster than a queue here.
		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if image[pos.y][pos.x] == color {
			image[pos.y][pos.x] = newColor

			if pos.x < len(image[pos.y])-1 && image[pos.y][pos.x+1] == color {
				queue = append(queue, position{x: pos.x + 1, y: pos.y})
			}

			if pos.x > 0 && image[pos.y][pos.x-1] == color {
				queue = append(queue, position{x: pos.x - 1, y: pos.y})
			}

			if pos.y < len(image)-1 && image[pos.y+1][pos.x] == color {
				queue = append(queue, position{x: pos.x, y: pos.y + 1})
			}

			if pos.y > 0 && image[pos.y-1][pos.x] == color {
				queue = append(queue, position{x: pos.x, y: pos.y - 1})
			}
		}
	}

	return image
}
