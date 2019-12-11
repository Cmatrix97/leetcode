/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package problems

import (
	"fmt"
	"math"
)

/*动态规划法
1.f(x) = max{f(x-1),x的从后向前(必须包含最后一项)最大子序和}
2.x的从后向前最大子序和=max{x-1的从后向前最大子序和,x-1的从后向前最大子序和+nums[x]}
*/
func maxSubArray1(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	dp := [][]int{}
	dp = append(dp, []int{nums[0], nums[0]})
	for i := 1; i < len(nums); i++ {
		dp2 := max(dp[i-1][1]+nums[i], nums[i])
		dp1 := max(dp[i-1][0], dp2)
		dp = append(dp, []int{dp1, dp2})
	}
	return dp[len(nums)-1][0]
}

/*动态规划法
优化写法
空间复杂度由O(n)降为O(1)
*/
func maxSubArray2(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}
		ans = int(math.Max(float64(ans), float64(sum)))
	}
	return ans
}

// @lc code=start

// @lc code=end

func Solve53() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Print(maxSubArray1(nums))
}
