package set_mismatch

// Returns the duplicate and missing numbers from nums. The slice ranges from 1..n, but there is one
// digit that has been duplicated, so there is also one missing.
//
// This solution uses cyclic sorting move all the numbers into their "correct" index positions
// (i.e. nums[i-1] == i) using constant memory. Then a second pass on the sorted array is done and
// any position that has the wrong number (i.e. nums[i] != i + 1) has both the duplicated and the
// missing number (nums[i] and i+1).
func findErrorNums(nums []int) []int {
	for i := range nums {
		for nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	errors := make([]int, 0, 2)
	for i := range nums {
		if nums[i] != i+1 {
			errors = append(errors, nums[i], i+1)
			break
		}
	}

	return errors
}
