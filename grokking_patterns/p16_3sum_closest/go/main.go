package _3sum_closest

import "fmt"
import "sort"

// Returns the sum of the three integers in the slice nums that are closest to the target.
//
// This solution uses a modified two pointers approach where one position is "fixed" as it passes
// over the sorted array and the other two positions use the opposite ends two pointer approach to
// check the optimal space of the remaining array after the "fixed" position. Since the array is
// sorted, the optimal move will be: if the sum is less then the target, move the lower pointer up
// and if the sum is higher then the target, move the upper pointer down. Maintaining the closest
// sum while checking the optimal sums allows for fixed memory usage.
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	// Since len(nums) <= 1000 and nums[i] <= 1000, then 10^6+1 is higher than any sum.
	closest := 1_000_001

	for c := 0; c < len(nums); c++ {
		for a, b := c+1, len(nums)-1; a < b; {
			sum := nums[a] + nums[b] + nums[c]

			if sum < target {
				a++
			} else {
				b--
			}

			if diff(sum, target) < diff(closest, target) {
				closest = sum
			}
		}
	}

	return closest
}

func diff(a, b int) int {
	if a < b {
		a, b = b, a
	}

	return a - b
}

//
// Bad approach of generating all combinations.
//

// Returns the sum of the three integers in the slice nums that are closest to the target.
func threeSumClosestBad(nums []int, target int) int {
	sums := checkCombos(nums, 3, 0, 0, target)

	fmt.Printf("nums %v sums %v\n", nums, sums)

	var closest *int
	for _, sum := range sums {
		fmt.Printf("closest is %d, comparing %d to %d, ", closest, sum, target)
		if closest == nil || diff(target, sum) < diff(target, *closest) {
			closest = &sum
		}
		fmt.Printf("closest is %d\n", closest)
	}

	return *closest
}

func checkCombos(nums []int, size, start, sum, target int) []int {
	if size == 0 {
		return []int{sum}
	}

	sums := []int{}

	for i := start; i < len(nums) && i < start+size; i++ {
		sum += nums[i]
		sums = append(sums, checkCombos(nums, size-1, i+1, sum, target)...)
	}

	return sums
}

// 5 -1 2 -1
// size = 3
// start = 0
// sum = 5
// target = 1
// result =
// i = 0 i <= 1
