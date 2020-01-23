/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package problems

import (
	"fmt"
	"math"
)

/*动态规划法
f(x) = max{f(x-1), x-min(x-1)}
*/
func maxProfit1_121(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(prices[i]-min)))
		min = int(math.Min(float64(min), float64(prices[i])))
	}
	return dp[len(prices)-1]
}

/*动态规划法
优化写法
*/
func maxProfit2_121(prices []int) int {
	minprice, maxprofit := math.MaxInt64, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxprofit {
			maxprofit = prices[i] - minprice
		}
	}
	return maxprofit
}

/*
股票问题通用解法
dp[i][k][0/1]
i:第i天(0开始)
k:至今最多进行k次交易
0/1:1表示持有股票
则边界:
dp[-1][k][0] = 0
dp[-1][k][1] = -INF
dp[i][0][0] = 0
dp[i][0][1] = -INF
状态转移方程:
dp[i][k][0] = max{dp[i-1][k][0], dp[i-1][k][1] + prices[i]}
dp[i][k][1] = max{dp[i-1][k][1], dp[i-1][k-1][0] - prices[i]}

121. k = 1
122. k = +INF
123. k = 2
188. k = k
309. k = +INF 买入时不考虑第i-1天卖出
714. k = +INF 卖出时将手续费减掉
*/
func maxProfit3_121(prices []int) int {
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
			dp[i][0], dp[i][1] = 0, -prices[0]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[N-1][0]
}

/*
模板优化
*/
func maxProfit4_121(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dpi0, dpi1 := 0, math.MinInt64
	for i := 0; i < len(prices); i++ {
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, -prices[i])
	}
	return dpi0
}

// @lc code=start

// @lc code=end

func Solve121() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1_121(prices))
}
