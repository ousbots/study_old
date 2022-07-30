package _3sum

import "sort"

func threeSum(nums []int) [][]int {
	sets := [][]int{}

	if len(nums) < 3 {
		return sets
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i += 1
		}

		low := i + 1
		high := len(nums) - 1

		for low < high {
			sum := nums[i] + nums[low] + nums[high]

			if sum == 0 {
				sets = append(sets, []int{nums[i], nums[low], nums[high]})
				low += 1
				high -= 1

				for low < len(nums) && nums[low] == nums[low-1] {
					low += 1
				}
				for high > 0 && nums[high] == nums[high+1] {
					high -= 1
				}
			} else if sum < 0 {
				low += 1
			} else {
				high -= 1
			}
		}
	}

	return sets
}
