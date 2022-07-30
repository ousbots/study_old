package p56_merge_intervals

import "sort"

func merge_slow(intervals [][]int) [][]int {
	output := [][]int{}

	for _, new := range intervals {
		temp := [][]int{}

		for _, old := range output {
			if new[1] < old[0] || new[0] > old[1] {
				temp = append(temp, old)
			}

			if new[0] >= old[0] && new[0] <= old[1] {
				new[0] = old[0]
			}

			if new[1] >= old[0] && new[1] <= old[1] {
				new[1] = old[1]
			}
		}

		temp = append(temp, new)
		output = temp
	}

	return output
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	keep := 0
	index := 1
	curr := intervals[0]

	for index < len(intervals) {
		next := intervals[index]
		index++

		if curr[1] < next[0] {
			intervals[keep] = curr
			keep++
			curr = next
			continue
		}

		if curr[1] < next[1] {
			curr[1] = next[1]
		}
	}

	intervals[keep] = curr
	keep++
	intervals = intervals[:keep]

	return intervals
}
