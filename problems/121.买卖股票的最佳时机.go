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
func maxProfit1(prices []int) int {
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

// @lc code=start
/*动态规划法
优化写法
*/
func maxProfit_121(prices []int) int {
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

// @lc code=end

func Solution121() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit_121(prices))
}
