package interview

import (
	"fmt"
	"math"
)

// 两个dp数组
func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums))
	maxTail := make([]int, len(nums))
	dp[0], maxTail[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxTail[i-1] > 0 {
			maxTail[i] = maxTail[i-1] + nums[i]
		} else {
			maxTail[i] = nums[i]
		}
		if maxTail[i] > dp[i-1] {
			dp[i] = maxTail[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}

// 降低空间复杂度
func maxSubArray2(nums []int) int {
	var sum int
	ans := math.MinInt64
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}

func SolveOffer42() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray2(nums))
}
