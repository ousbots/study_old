package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	total := 1
	zeros := 0

	for _, num := range nums {
		if num == 0 {
			zeros += 1
		} else {
			total *= num
		}
	}

	for i := range nums {
		if zeros > 1 {
			nums[i] = 0
		} else if zeros == 1 {
			if nums[i] == 0 {
				nums[i] = total
			} else {
				nums[i] = 0
			}
		} else {
			nums[i] = total / nums[i]
		}
	}

	return nums
}
