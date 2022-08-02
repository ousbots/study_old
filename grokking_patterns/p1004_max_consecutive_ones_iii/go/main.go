package max_consecutive_ones_iii

// Returns the maximum number of consecutive 1's in the nums slice if at most k 0's can be flipped.
// The elements in the nums slice are either 0 or 1.
func longestOnes(nums []int, k int) int {
	flipped, back, max := 0, 0, 0
	positions := []int{}

	for front, value := range nums {
		if value == 0 {
			positions = append(positions, front)
			flipped++

			for flipped > k && back <= front {
				if back == positions[0] {
					flipped--
					positions = positions[1:]
				}
				back++
			}
		}

		if front-back >= max {
			max = front - back + 1
		}
	}

	return max
}
