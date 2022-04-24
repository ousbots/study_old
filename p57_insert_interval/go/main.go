package insert_interval

func bSearch(intervals [][]int, val int) int {
	if len(intervals) <= 1 {
		return 0
	}

	low := 0
	high := len(intervals) - 1
	middle := 0

	for low <= high {
		middle = (low + high) / 2

		if intervals[middle][0] > val {
			high = middle - 1
		} else if intervals[middle][1] < val {
			low = middle + 1
		} else {
			break
		}
	}

	return middle
}

func insert_binary_search(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	front := bSearch(intervals, newInterval[0])
	back := bSearch(intervals, newInterval[1])

	if front == back {
		if intervals[front][1] < newInterval[0] {
			intervals = append(intervals[:front+1], intervals[front:]...)
			intervals[front+1] = newInterval

			return intervals
		}

		if intervals[front][0] > newInterval[1] {
			intervals = append(intervals[:front+1], intervals[front:]...)
			intervals[front] = newInterval

			return intervals
		}

		if intervals[front][1] < newInterval[1] {
			intervals[front][1] = newInterval[1]
		}
	} else {
		if intervals[front][1] < newInterval[0] {
			front += 1
		}

		if intervals[back][0] > newInterval[1] {
			back -= 1
		}

		if intervals[back][1] < newInterval[1] {
			intervals[front][1] = newInterval[1]
		} else {
			intervals[front][1] = intervals[back][1]
		}

		intervals = append(intervals[:front+1], intervals[back+1:]...)
	}

	if intervals[front][0] > newInterval[0] {
		intervals[front][0] = newInterval[0]
	}

	return intervals
}

func insert_linear(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	front := 0
	for front < len(intervals) && intervals[front][1] < newInterval[0] {
		front += 1
	}

	if front == len(intervals) && intervals[front-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}

	back := front
	for back < len(intervals)-1 && intervals[back+1][0] <= newInterval[1] {
		back += 1
	}

	if back == 0 && intervals[back][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}

	if back == front && intervals[front][0] > newInterval[1] {
		intervals = append(intervals[:front+1], intervals[front:]...)
		intervals[front] = newInterval
		return intervals
	}

	if intervals[front][0] > newInterval[0] {
		intervals[front][0] = newInterval[0]
	}

	if intervals[back][1] < newInterval[1] {
		intervals[back][1] = newInterval[1]
	}

	if front != back {
		intervals[front][1] = intervals[back][1]
		intervals = append(intervals[:front+1], intervals[back+1:]...)
	}

	return intervals
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	front := bSearch(intervals, newInterval[0])
	if intervals[front][1] < newInterval[0] {
		front += 1
	}

	if front == len(intervals) && intervals[front-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}

	back := bSearch(intervals, newInterval[1])
	if intervals[back][0] > newInterval[1] {
		back -= 1
	}

	if back == 0 && intervals[back][0] > newInterval[1] {
		return append([][]int{newInterval}, intervals...)
	}

	if back < front {
		intervals = append(intervals[:front+1], intervals[front:]...)
		intervals[front] = newInterval
		return intervals
	}

	if intervals[front][0] > newInterval[0] {
		intervals[front][0] = newInterval[0]
	}

	if intervals[back][1] < newInterval[1] {
		intervals[back][1] = newInterval[1]
	}

	if front != back {
		intervals[front][1] = intervals[back][1]
		intervals = append(intervals[:front+1], intervals[back+1:]...)
	}

	return intervals
}
