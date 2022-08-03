package sort_colors

// In-place sorts the nums slice so that 0s are first, 1s are second, and 2s are last.
//
// This solution uses two pointers, one for the back (0s) and one for the front (2s). The nums
// slice is iterated through and if nums[i] is 0, then it is swapped with the back position and
// the back pointer incremented and if nums[i] is 2 and i < front, then it is swapped with the front
// position and the front pointer decremented.
func sortColors(nums []int) {
	back := 0
	front := len(nums) - 1

	for i := 0; i <= front; i++ {
		if nums[i] == 0 && i != back {
			nums[back], nums[i] = nums[i], nums[back]
			back++
			i--
		} else if nums[i] == 2 && i != front {
			nums[front], nums[i] = nums[i], nums[front]
			front--
			i--
		}
	}
}
