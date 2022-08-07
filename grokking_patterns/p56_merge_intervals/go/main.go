package merge_interval

import "sort"

// Returns all the intervals with overlapping intervals merged. The slice is not sorted, but each
// individual interval is (i.e. for intervals[i] intervals[i][0] <= intervals[i][1]).
//
// This solution first sorts the intervals slice by the first element in each element, then
// iterates over the sorted intervals checking if the current interval overlaps with the current.
// If they do overlap, they are merged and added to the slice of merged intervals. If not, the
// current interval is added as-is.
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
