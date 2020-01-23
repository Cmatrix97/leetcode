/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */
package problems

import (
	"fmt"
	"math"
)

/*
一次交易手续费为fee
每次卖出（或买入）减去手续费
*/
func maxProfit1_714(prices []int, fee int) int {
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
			dp[i][0], dp[i][1] = 0, -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[N-1][0]
}

/*
注意溢出
dpi1设为MinInt32
*/
func maxProfit2_714(prices []int, fee int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		temp := dpi0
		dpi0 = max(dpi0, dpi1+prices[i]-fee)
		dpi1 = max(dpi1, temp-prices[i])
	}
	return dpi0
}

// @lc code=start

// @lc code=end
func Solve714() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println(maxProfit1_714(prices, fee))
}
