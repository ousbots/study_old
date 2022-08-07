package interval_list_intersections

// Returns a slice of the overlapping intervals from the two slices of intervals. Each set of
// intervals is disjoint and ordered.
//
// This solution keep track of an index for each list and for each iteration, checks the intervals
// at each index to see if there is any overlap and stores it if there is. Whichever list's
// current interval's end is less then the other interval's end is the one to be incremented.
func intervalIntersection(firstList, secondList [][]int) [][]int {
	overlaps := [][]int{}

	first, second, fEnd, sEnd := 0, 0, len(firstList), len(secondList)

	for first < fEnd && second < sEnd {
		front := firstList[first][0]
		if secondList[second][0] > front {
			front = secondList[second][0]
		}

		back := firstList[first][1]
		if secondList[second][1] < back {
			back = secondList[second][1]
		}

		if front <= back {
			overlaps = append(overlaps, []int{front, back})
		}

		if firstList[first][1] < secondList[second][1] {
			first++
		} else {
			second++
		}
	}

	return overlaps
}

// Returns a slice of the overlapping intervals from the two slices of intervals. Each set of
// intervals is disjoint and ordered.
//
// This solution iterates through the intervals in one of the slices and checks if it overlaps
// with the intervals of the other list. The interval to check is determined by skipping intervals
// until the current interval's start point is not greater than the intervals from the other list's
// end points. If the two intervals overlap, the overlapping interval is added to the slice of
// intervals to return.
func intervalIntersectionSlow(firstList, secondList [][]int) [][]int {
	overlaps := [][]int{}
	sIndex := 0

	for _, firstInterval := range firstList {
		for _, secondInterval := range secondList[sIndex:] {
			if firstInterval[0] <= secondInterval[1] && firstInterval[1] >= secondInterval[0] {
				temp := overlap(firstInterval, secondInterval)
				if len(temp) > 0 {
					overlaps = append(overlaps, temp)
				}

				if firstInterval[0] > secondInterval[1] {
					sIndex++
				}
			}
		}
	}

	return overlaps
}

func overlap(a, b []int) []int {
	start := a[0]
	if b[0] > start {
		start = b[0]
	}

	end := a[1]
	if b[1] < end {
		end = b[1]
	}

	if start <= end {
		return []int{start, end}
	} else {
		return nil
	}
}

// [1 5] [5 6]
// start = 5
// end 5