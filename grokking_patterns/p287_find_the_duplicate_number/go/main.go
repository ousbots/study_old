package find_the_duplicate_number

// Return the duplicate integer in nums. The nums slice contains numbers 1..n and has a size of
// n+1. Needs to be solved without modifying nums and using constant space.
//
// This solution uses the fast & slow pointers to find the start of the cycle. Because the integers
// in nums range from 1..n with one duplicate, this means each integer points to a position in the
// slice. And the duplicate will create a cycle. I'm not exactly sure why, but the duplicate is also
// always the start of the cycle. So finding the start of the cycle also finds the duplicate.
func findDuplicate(nums []int) int {
	slow := 0
	fast := 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0

	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
