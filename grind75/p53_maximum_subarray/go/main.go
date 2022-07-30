package maximum_subarray

func maxSubArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	max := nums[0]
	total := 0

	for _, num := range nums {
		total += num

		if total > 0 {
			if total > max {
				max = total
			}
		} else {
			total = 0

			if num > max {
				max = num
			}
		}
	}

	return max
}
