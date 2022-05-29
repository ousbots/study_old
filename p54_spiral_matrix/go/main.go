package p54_spiral_matrix

func spiralOrder(matrix [][]int) []int {
	ymin := 0
	ymax := len(matrix) - 1
	xmin := 0
	xmax := len(matrix[0]) - 1

	output := []int{}

	vert := false
	inc := 1

	length := len(matrix) * len(matrix[0])

	x, y := 0, 0
	for i := 0; i < length; i++ {
		output = append(output, matrix[y][x])

		if vert {
			y += inc

			if y > ymax {
				y = ymax
				xmax--
				inc = -1
				x += inc
				vert = false
			} else if y < ymin {
				y = ymin
				xmin++
				inc = 1
				x += inc
				vert = false
			}
		} else {
			x += inc

			if x > xmax {
				x = xmax
				ymin++
				inc = 1
				y += inc
				vert = true
			} else if x < xmin {
				x = xmin
				ymax--
				inc = -1
				y += inc
				vert = true
			}
		}
	}

	return output
}
