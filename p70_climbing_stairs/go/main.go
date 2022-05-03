package climbing_stairs

func climbStairs(n int) int {
	values := make([]int, n+1)
	values[0] = 1
	values[1] = 1

	for i := 2; i <= n; i++ {
		values[i] = values[i-1] + values[i-2]
	}

	return values[n]
}
