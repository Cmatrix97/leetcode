/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */
package problems

import "fmt"

// @lc code=start
func maxProfit_122(prices []int) int {
	maxprofit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxprofit += prices[i] - prices[i-1]
		}
	}
	return maxprofit
}

// @lc code=end

func Solution122() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit_122(prices))
}
