package p78_subsets

func subsets(nums []int) [][]int {
	return generate(nums, []int{}, 0)
}

func generate(nums, set []int, index int) [][]int {
	if index >= len(nums) {
		return [][]int{set}
	}

	b := make([]int, len(set))
	copy(b, set)
	b = append(b, nums[index])

	return append(generate(nums, b, index+1), generate(nums, set, index+1)...)
}
