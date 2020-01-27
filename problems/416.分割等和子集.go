/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */
package problems

import (
	"fmt"
	"math"
)

/*
0-1背包
*/
func canPartition1(nums []int) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	dp := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		dp[i] = math.MinInt64
	}
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			if dp[j] == sum {
				return true
			}
		}
	}
	return false
}

/*
int --> bool
dp[i]表示当前是否能凑出和为i
*/
func canPartition2(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	if sum&1 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			if dp[sum] {
				return true
			}
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return false
}

// @lc code=start

// @lc code=end
func Solve416() {
	nums := []int{1, 2, 3, 5}
	fmt.Println(canPartition1(nums))
}
