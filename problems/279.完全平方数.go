/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */
package problems

import (
	"fmt"
	"math"
)

/*
n=270超时
*/
func numSquares1(n int) int {
	var nums []int
	for i := 1; i*i <= n; i++ {
		nums = append(nums, i*i)
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
	min := math.MaxInt64
	var recurse func(count, sum, start int)
	recurse = func(count, sum, start int) {
		if sum == n {
			min = int(math.Min(float64(min), float64(count)))
			return
		}
		for i := start; i < len(nums); i++ {
			if sum+nums[i] > n {
				continue
			}
			recurse(count+1, sum+nums[i], i)
		}
	}
	recurse(0, 0, 0)
	return min
}

/*
动态规划
dp[i] = min(dp[i-所有平方数]) + 1
*/
func numSquares2(n int) int {
	dp := make([]int, n+1)
	m := make(map[int]bool)
	for i := 1; i*i <= n; i++ {
		m[i*i], dp[i*i] = true, 1
	}
	for i := 2; i <= n; i++ {
		if dp[i] == 1 {
			continue
		}
		min := math.MaxInt64
		for k := range m {
			if i < k {
				continue
			}
			if dp[i-k] < min {
				min = dp[i-k]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}

/*
dp简化
*/
func numSquares3(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			if dp[i-j*j]+1 < dp[i] {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

/*
拉格朗日四平方数和定理
任何一个正整数都可以表示成不超过四个整数的平方之和。
推论：满足四数平方和定理的数n（四个整数的情况），必定满足 n=4^a(8b+7)
*/
func numSquares4(n int) int {
	//先根据上面提到的公式来缩小n
	for n%4 == 0 {
		n /= 4
	}
	//如果满足公式 则返回4
	if n%8 == 7 {
		return 4
	}
	//在判断缩小后的数是否可以由一个数的平方或者两个数平方的和组成
	a := 0
	for (a * a) <= n {
		b := int(math.Sqrt(float64(n - a*a)))
		if a*a+b*b == n {
			//如果可以 在这里返回
			if a != 0 && b != 0 {
				return 2
			} else {
				return 1
			}
		}
		a++
	}
	//如果不行 返回3
	return 3
}

// @lc code=start

// @lc code=end
func Solve279() {
	n := 13
	fmt.Println(numSquares1(n))
}
