/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package problems

import (
	"fmt"
	"math"
)

/*
冷冻期为1天
第i天选择buy，要从第i-2天sell状态转移
*/
func maxProfit1_309(prices []int) int {
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
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < N; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		if i == 1 {
			dp[i][1] = max(dp[i-1][1], -prices[i])
			continue
		}
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[N-1][0]
}

/*
临时变量存储dp_i-2_0
*/
func maxProfit2_309(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dppre0, dpi0, dpi1 := 0, 0, math.MinInt64
	for i := 0; i < len(prices); i++ {
		temp := dpi0
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, dppre0-prices[i])
		dppre0 = temp
	}
	return dpi0
}

// @lc code=start

// @lc code=end
func Solve309() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit1_309(prices))
}
