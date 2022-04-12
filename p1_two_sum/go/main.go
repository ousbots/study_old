package two_sum

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, val := range nums {
		if j, exists := seen[target-val]; exists {
			return []int{i, j}
		}

		seen[val] = i
	}

	return nil
}
