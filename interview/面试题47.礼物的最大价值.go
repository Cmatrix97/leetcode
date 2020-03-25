package interview

import "fmt"

// 递归+备忘录
func maxValue1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	m := make(map[int]int)
	var recurse func(i, j int) int
	recurse = func(i, j int) int {
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		if v, ok := m[i*N+j]; ok {
			return v
		}
		var res int
		if i == 0 {
			res = recurse(i, j-1) + grid[i][j]
		} else if j == 0 {
			res = recurse(i-1, j) + grid[i][j]
		} else {
			res = max(recurse(i, j-1), recurse(i-1, j)) + grid[i][j]
		}
		m[i*N+j] = res
		return res
	}
	return recurse(M-1, N-1)
}

// 二维dp
func maxValue2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	dp := make([][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([]int, N)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < M; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < N; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[M-1][N-1]
}

// 一维dp
func maxValue3(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	dp := make([]int, N+1)
	max := func (x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			dp[j+1] = max(dp[j], dp[j+1]) + grid[i][j]
		}
	}
	return dp[N]
}

func SolveOffer47() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue3(grid))
}
