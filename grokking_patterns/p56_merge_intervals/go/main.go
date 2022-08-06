package merge_interval

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	merged := [][]int{intervals[0]}
	intervals = intervals[1:]
	mIndex := 0

	for i := range intervals {
		if !overlap(merged[mIndex], intervals[i]) {
			merged = append(merged, intervals[i])
			mIndex++
		} else {
			start := merged[mIndex][0]
			if intervals[i][0] < start {
				start = intervals[i][0]
			}

			end := intervals[i][1]
			if merged[mIndex][1] > end {
				end = merged[mIndex][1]
			}

			merged[mIndex][0] = start
			merged[mIndex][1] = end
		}
	}

	return merged
}

func overlap(a, b []int) bool {
	start := a[0]
	if b[0] > start {
		start = b[0]
	}

	end := a[1]
	if b[1] < end {
		start = b[1]
	}

	return end-start >= 0
}
