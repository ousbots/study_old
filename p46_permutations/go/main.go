package permutations

func permute(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, &result, 0)

	return result
}

func backtrack(nums []int, perms *[][]int, start int) {
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*perms = append(*perms, tmp)
	}

	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrack(nums, perms, start+1)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
