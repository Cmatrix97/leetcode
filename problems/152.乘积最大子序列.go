/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
package problems

import (
	"fmt"
	"math"
)

/*
动态规划
dp[i]表示前i个数的最大乘积
max[i]表示以第i个数结尾的最大乘积
min[i]表示以第i个数结尾的最大乘积
dp[i] = max{dp[i-1], max[i]}
(1)如果nums[i]为正数, max[i] = max{max[i-1]*nums[i], nums[i]}
(1)如果nums[i]为负数, max[i] = max{min[i-1]*nums[i], nums[i]}
min的求法与max相反
*/
func maxProduct1(nums []int) int {
	minSuffix, maxSuffix, dp := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	minSuffix[0], maxSuffix[0], dp[0] = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			minSuffix[i] = int(math.Min(float64(nums[i]), float64(nums[i]*minSuffix[i-1])))
			maxSuffix[i] = int(math.Max(float64(nums[i]), float64(nums[i]*maxSuffix[i-1])))
		} else if nums[i] < 0 {
			minSuffix[i] = int(math.Min(float64(nums[i]), float64(nums[i]*maxSuffix[i-1])))
			maxSuffix[i] = int(math.Max(float64(nums[i]), float64(nums[i]*minSuffix[i-1])))
		}
		dp[i] = int(math.Max(float64(dp[i-1]), float64(maxSuffix[i])))
	}
	return dp[len(nums)-1]
}

/*
降低空间复杂度为O(1)
*/
func maxProduct2(nums []int) int {
	max, imax, imin := math.MinInt64, 1, 1
	for i := range nums {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = int(math.Max(float64(nums[i]), float64(nums[i]*imax)))
		imin = int(math.Min(float64(nums[i]), float64(nums[i]*imin)))
		max = int(math.Max(float64(max), float64(imax)))
	}
	return max
}

// @lc code=start

// @lc code=end

func Solve152() {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct1(nums))
}
