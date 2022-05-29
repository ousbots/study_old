package p416_partition_equal_subset_sum

import "sort"

// Dynamic programming solution is much faster O(n)
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	dp := make([]bool, sum+1)
	dp[0] = true

	for _, num := range nums {
		for j := sum; j >= 0; j-- {
			if dp[j] {
				dp[j+num] = true
			}
		}
	}

	return dp[sum/2]
}

func canPartitionDFS(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}

	sort.Ints(nums)

	return find(nums, sum/2, len(nums)-1, 0, 0)
}

func find(a []int, check, index, as, bs int) bool {
	if as > check || bs > check {
		return false
	}

	if as == check {
		return true
	}

	if index < 0 {
		return false
	}

	return find(a, check, index-1, as+a[index], bs) || find(a, check, index-1, as, bs+a[index])
}
