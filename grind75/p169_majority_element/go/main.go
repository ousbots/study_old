package majority_element

func majorityElementSlow(nums []int) int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	majority := 0
	max := 0

	for key, val := range count {
		if val > max {
			majority = key
			max = val
		}
	}

	return majority
}

func majorityElement(nums []int) int {
	majority := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			majority = num
			count++
		} else if num == majority {
			count++
		} else {
			count--
		}
	}

	return majority
}
