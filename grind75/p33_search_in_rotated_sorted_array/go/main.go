package search_in_rotated_sorted_array

func search(num []int, target int) int {
	lo := 0
	hi := len(num) - 1

	for lo <= hi {
		middle := (lo + hi) / 2

		if target < num[middle] {
			if num[hi] < num[middle] && num[lo] > target {
				lo = middle + 1
			} else {
				hi = middle - 1
			}
		} else if target > num[middle] {
			if num[lo] > num[middle] && num[hi] < target {
				hi = middle - 1
			} else {
				lo = middle + 1
			}
		} else {
			return middle
		}
	}

	return -1
}
