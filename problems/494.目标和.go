/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
package problems

import (
	"fmt"
)

/*
dfs暴力穷举
*/
func findTargetSumWays1(nums []int, S int) int {
	var count int
	var dfs func(index int, sum int)
	dfs = func(index int, sum int) {
		if index == len(nums) {
			if sum == S {
				count++
			}
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return count
}

// @lc code=start
/*
转换为0-1背包
{+}-{-} = S
{+}+{-} = sum
{-} = (sum-S)/2
即问题转化为找一些元素使其和为(sum-S)/2
dp[j]表示当前能组成和为j的方法总数
*/
func findTargetSumWays2(nums []int, S int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum < S || (sum-S)&1 == 1 {
		return 0
	}
	target := (sum - S) / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}

// @lc code=end
func Solve494() {
	nums := []int{1, 1, 1, 1, 1}
	S := 3
	fmt.Println(findTargetSumWays1(nums, S))
}
