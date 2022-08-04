package subarray_product_less_than_k

// Return the number of contiguous subarrays in nums whose product is less than K.
//
// This solution use a two pointer pattern where the back pointer increments over the nums slice
// and for every increment, the forward pointer increments until the product is larger than k.
func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0

	for back := range nums {
		product := nums[back]
		if product < k {
			count++
		}

		for front := back + 1; front < len(nums) && product < k; front++ {
			product *= nums[front]
			if product < k {
				count++
			}
		}
	}

	return count
}

// Return the number of contiguous subarrays in nums whose product is less than K.
//
// This solution attempted to move the front pointer forward and follow it with the back pointer
// when the product is too large, but that only captures the largest subarrays that are less than
// k. This could be fixed by calculating the number of permutations, but there are simpler ways
// of solving this.
func numSubarrayProductLessThanKBad(nums []int, k int) int {
	count, back := 0, 0
	product := 1
	for front := range nums {
		product *= nums[front]

		if product < k {
			count++
		} else {
			for product > k && back < front {
				product /= nums[back]
				back++
			}

			if product < k {
				count++
			}
		}
	}

	for back < len(nums) {
		product /= nums[back]

		if product < k {
			count++
		}

		back++
	}

	return count
}
