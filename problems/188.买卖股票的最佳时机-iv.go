/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */
package problems

import "fmt"

/*
k > N / 2时视为k = +INF
*/
func maxProfit1_188(k int, prices []int) int {
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
	maxProfitKINF := func() int {
		dp := make([][]int, N)
		for i := 0; i < N; i++ {
			dp[i] = make([]int, 2)
		}
		dpi0, dpi1 := 0, -prices[0]
		for i := 0; i < N; i++ {
			dpi0 = max(dpi0, dpi1+prices[i])
			dpi1 = max(dpi1, dpi0-prices[i])
		}
		return dpi0
	}
	if k > N/2 {
		return maxProfitKINF()
	}
	dp := make([][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < N; i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[0]
				continue
			}
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[N-1][k][0]
}

// @lc code=start

// @lc code=end
func Solve188() {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Println(maxProfit1_188(k, prices))
}
