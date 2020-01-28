/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */
package problems

import "fmt"

// @lc code=start
/*
完全背包
*/
func change1(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// @lc code=end
func Solve518() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change1(amount, coins))
}
