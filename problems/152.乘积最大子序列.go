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


// 动态规划
// ans表示前i个数的最大乘积
// max[i]表示以第i个数结尾的最大乘积
// min[i]表示以第i个数结尾的最小乘积
// ans = max{ans, max[i]}
// (1)如果nums[i]为正数, max[i] = max{max[i-1]*nums[i], nums[i]}
// (1)如果nums[i]为负数, max[i] = max{min[i-1]*nums[i], nums[i]}
// min的求法与max相反
func maxProduct1(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	maxArr, minArr := make([]int, len(nums)), make([]int, len(nums))
	maxArr[0], minArr[0] = nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= 0 {
			maxArr[i] = max(nums[i], maxArr[i-1]*nums[i])
			minArr[i] = min(nums[i], minArr[i-1]*nums[i])
		} else {
			maxArr[i] = max(nums[i], minArr[i-1]*nums[i])
			minArr[i] = min(nums[i], maxArr[i-1]*nums[i])
		}
		ans = max(ans, maxArr[i])
	}
	return ans
}

// 降低空间复杂度为O(1)
func maxProduct2(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	ans, imax, imin := math.MinInt64, 1, 1
	for _, num := range nums {
		if num < 0 {
			imax, imin = imin, imax
		}
		imax = max(num, imax*num)
		imin = min(num, imin*num)
		ans = max(ans, imax)
	}
	return ans
}

// @lc code=start

// @lc code=end

func Solve152() {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct1(nums))
}
