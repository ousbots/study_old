package main

import (
	"sort"
)

func countRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int64, len(nums) + 1)
	sums[0] = 0

	for i, num := range nums {
		sums[i+1] = sums[i] + int64(num)
	}
	
	return countRangeMerge(sums, 0, len(sums), lower, upper)
}

func countRangeMerge(sums []int64, start, end, lower, upper int) int {
	if end - start <= 1 {
		return 0
	}
	
	mid := (start + end) / 2
	
	count := countRangeMerge(sums, start, mid, lower, upper) + countRangeMerge(sums, mid, end, lower, upper)
	
	j, k := mid, mid
	
	for _, val := range sums[start:mid] {
		for j < end && sums[j] - val < int64(lower) {
			j++
		}
		
		for k < end && sums[k] - val <= int64(upper) {
			k++
		}
		
		count += k - j
	}
	
	sort.Slice(sums[start:end], func(i, j int) bool { return sums[start:end][i] < sums[start:end][j] })

	return count
}
