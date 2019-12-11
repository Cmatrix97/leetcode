/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package problems

import (
	"fmt"
	"math"
)

/*动态规划
f(x) = max{f(x-1), x + f(x-2)}
*/
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(nums[i]+dp[i-2])))
	}
	return dp[len(nums)-1]
}

/*
简化写法，空间复杂度O(n) --> O(1)
*/
func rob2(nums []int) int {
	var max, pre, temp float64
	for _, v := range nums {
		temp = max
		max = math.Max(max, float64(v)+pre)
		pre = temp
	}
	return int(max)
}

// @lc code=start

// @lc code=end

func Solve198() {
	nums := []int{}
	// nums := []int{1, 2, 3, 1}
	fmt.Println(rob1(nums))
}
