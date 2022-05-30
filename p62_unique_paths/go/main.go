package p62_unique_paths

func uniquePathsSlow(m, n int) int {
	return find(m, n, 0, 0)
}

func find(m, n, y, x int) int {
	if y >= m || x >= n {
		return 0
	}

	if y == m-1 && x == n-1 {
		return 1
	}

	return find(m, n, y+1, x) + find(m, n, y, x+1)
}

func uniquePathsFib(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m == 2 {
		return n
	}

	if n == 2 {
		return m
	}

	return uniquePathsFib(m-1, n) + uniquePathsFib(m, n-1)
}

// Dynamic programming algorithm: build the total number of moves to reach the goal.
func uniquePaths(m, n int) int {
	options := make([][]int, m)

	for i := range options {
		options[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		options[i][0] = 1
	}

	for i := 0; i < n; i++ {
		options[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			options[i][j] = options[i-1][j] + options[i][j-1]
		}
	}

	return options[m-1][n-1]
}
