package find_all_duplicates_in_an_array

// Returns a slice of all integers in the nums slice that appear twice. The integers in the nums
// slice range from 1 to n and occur either once or twice.
//
// This solution uses a slightly modified cycle sort, since the numbers range from 1 to n, the
// position of the number does not need to be calculated. Instead, if a number is in the wrong
// place then it is swapped into the correct place. The swapping continues until the current
// position has the correct number. Once the slice has been iterated and swapped through, the
// duplicate numbers are in the missing numbers positions (1 <= nums[i] <= n, so a duplicate means
// a number in the sequence is missing). So the slice needs to be iterated over one more time,
// moving numbers in incorrect positions to the front then returning the slice of numbers moved to
// the front.
func findDuplicates(nums []int) []int {
	for start := range nums {
		for nums[nums[start]-1] != nums[start] {
			nums[nums[start]-1], nums[start] = nums[start], nums[nums[start]-1]
		}
	}

	count := 0
	for i := range nums {
		if nums[i] != i+1 {
			nums[count] = nums[i]
			count++
		}
	}

	return nums[:count]
}