package binary_search

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := (low + high) / 2

		if nums[middle] < target {
			low = middle + 1
			continue
		} else if nums[middle] > target {
			high = middle - 1
			continue
		} else {
			return middle
		}
	}

	return -1
}
