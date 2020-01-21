/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */
package problems

import (
	"fmt"
	"math"
)

/*
暴力法
*/
func maximalSquare1(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	M, N := len(matrix), len(matrix[0])
	getSquare := func(m, n int) int {
		x := 1
	flag:
		for m+x <= M && n+x <= N {
			for i := m; i < m+x; i++ {
				for j := n; j < n+x; j++ {
					if matrix[i][j] == '0' {
						break flag
					}
				}
			}
			x++
		}
		return x - 1
	}
	var max int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == '1' {
				max = int(math.Max(float64(max), float64(getSquare(i, j))))
			}
		}
	}
	return max * max
}

/*
动态规划
dp[i][j]表示包含(i,j)的最大正方形边长
dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
*/
func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	M, N := len(matrix), len(matrix[0])
	dp := make([][]int, M+1)
	for i := 0; i < M+1; i++ {
		dp[i] = make([]int, N+1)
	}
	var max int
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1]))) + 1
				max = int(math.Max(float64(max), float64(dp[i][j])))
			}
		}
	}
	return max * max
}

/*
可将二维dp用一维dp+prev替换
prev     dp[j]
dp[j-1]  cur
*/
func maximalSquare3(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	M, N := len(matrix), len(matrix[0])
	dp := make([]int, N+1)
	var max, prev int
	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			temp := dp[j]
			if matrix[i-1][j-1] == '1' {
				dp[j] = int(math.Min(math.Min(float64(dp[j-1]), float64(prev)), float64(dp[j]))) + 1
				max = int(math.Max(float64(max), float64(dp[j])))
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return max * max
}

// @lc code=start

// @lc code=end
func Solve221() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare1(matrix))
}
