/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package problems

import (
	"fmt"
	"math"
)

/*
k = +INF
*/
func maxProfit1_122(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}

func maxProfit2_122(prices []int) int {
	N := len(prices)
	if N == 0 {
		return 0
	}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, 2)
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < N; i++ {
		if i == 0 {
			dp[i][0], dp[i][1] = 0, -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[N-1][0]
}

func maxProfit3_122(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dpi0, dpi1 := 0, math.MinInt64
	for i := 0; i < len(prices); i++ {
		temp := dpi0
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, temp-prices[i])
	}
	return dpi0
}

// @lc code=start

// @lc code=end

func Solve122() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1_122(prices))
}
