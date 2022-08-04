package _4sum

import "sort"

// Returns an array of all sets of four elements from nums that add up to the target. Each set
// should be unique.
//
// This is a modified two pointers solution, that uses two "fixed" points and front and back
// pointers to search for optimal sets. If the sum is less than the target, the back pointer should
// be incremented, but if the sum is greater than the target, the front pointer should be
// decremented. The elements of nums are not unique, so one of the pointers should be changed when
// the sum equals the target.
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	if len(nums) < 4 {
		return [][]int{}
	}

	seen := make(map[[4]int]struct{})
	tuples := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		fixed := nums[i]
		temp := nums[i+1:]

		for j := 0; j < len(temp)-2; j++ {
			back := j + 1
			front := len(temp) - 1

			for back < front {
				sum := fixed + temp[j] + temp[back] + temp[front]

				if sum < target {
					back++
				} else if sum > target {
					front--
				} else {
					set := [4]int{fixed, temp[j], temp[back], temp[front]}
					if _, exists := seen[set]; !exists {
						tuples = append(tuples, set[:])
						seen[set] = struct{}{}
					}
					back++
				}
			}
		}
	}

	return tuples
}
