package insert_interval

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	merged := [][]int{intervals[0]}
	mIndex := 0

	for _, interval := range intervals {
		if !overlap(merged[mIndex], interval) {
			merged = append(merged, interval)
			mIndex++
		} else {
			start := interval[0]
			if merged[mIndex][0] < start {
				start = merged[mIndex][0]
			}

			end := interval[1]
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

	end := b[1]
	if a[1] < end {
		end = a[1]
	}

	return end-start >= 0
}
