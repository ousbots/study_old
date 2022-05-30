package p11_container_with_most_water

func maxArea(height []int) int {
	first, second := 0, len(height)-1
	max := 0

	for second > first {
		min := height[first]
		if height[second] < min {
			min = height[second]
		}

		area := (second - first) * min
		if area > max {
			max = area
		}

		if height[first] < height[second] {
			first++
		} else {
			second--
		}
	}

	return max
}
