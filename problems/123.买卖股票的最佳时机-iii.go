/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */
package problems

import "fmt"

/*
k = 2
*/
func maxProfit1_123(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	N := len(prices)
	if N == 0 {
		return 0
	}
	maxK := 2
	dp := make([][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, maxK+1)
		for j := 0; j < maxK+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < N; i++ {
		for k := maxK; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[N-1][maxK][0]
}

// @lc code=start

// @lc code=end
func Solve123() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit1_123(prices))
}
