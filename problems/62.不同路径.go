/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */
package problems

import "fmt"

/*
简单动态规划
*/
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
		for j := 0; j < n; j++ {
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=start

// @lc code=end

func Solve62() {
	m, n := 3, 2
	fmt.Println(uniquePaths1(m, n))
}
