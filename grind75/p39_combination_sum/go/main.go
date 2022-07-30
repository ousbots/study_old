package combination_sum

func combinationSumCopy(candidates []int, target int) [][]int {
	return find(candidates, target, 0, []int{})
}

func find(vals []int, total int, index int, sum []int) [][]int {
	if total == 0 {
		tmp := make([]int, len(sum))
		copy(tmp, sum)
		return [][]int{tmp}
	} else if total < 0 || index == len(vals) {
		return nil
	} else {
		sums := [][]int{}
		sums = append(sums, find(vals, total, index+1, sum)...)

		sum = append(sum, vals[index])
		sums = append(sums, find(vals, total-vals[index], index, sum)...)

		return sums
	}
}

func combinationSumFor(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrackFor(&result, []int{}, candidates, 0, target)

	return result
}

func backtrackFor(result *[][]int, sum, vals []int, index, total int) {
	if total < 0 || index >= len(vals) {
		return
	}

	if total == 0 {
		tmp := make([]int, len(sum))
		copy(tmp, sum)
		*result = append(*result, tmp)
		return
	}

	for i := index; i < len(vals); i++ {
		sum = append(sum, vals[i])
		backtrackFor(result, sum, vals, i, total-vals[i])
		sum = sum[:len(sum)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrack(&result, []int{}, candidates, 0, target)

	return result
}

func backtrack(result *[][]int, sum, vals []int, index, total int) {
	if total < 0 || index >= len(vals) {
		return
	}

	if total == 0 {
		tmp := make([]int, len(sum))
		copy(tmp, sum)
		*result = append(*result, tmp)
		return
	}

	backtrack(result, sum, vals, index+1, total)

	sum = append(sum, vals[index])
	backtrack(result, sum, vals, index, total-vals[index])
}
