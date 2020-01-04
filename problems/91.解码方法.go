/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */
package problems

import (
	"fmt"
	"strconv"
)

/*
动态规划
*/
func numDecodings1(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if s[i-2] != '0' {
			if v, _ := strconv.Atoi(s[i-2 : i]); v >= 1 && v <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n]
}

/*
空间复杂度降为O(1)
*/
func numDecodings2(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	t0, t1, t2 := 1, 1, 1
	for i := 2; i <= n; i++ {
		t2 = 0
		if s[i-1] != '0' {
			t2 += t1
		}
		if s[i-2] != '0' {
			if v, _ := strconv.Atoi(s[i-2 : i]); v >= 1 && v <= 26 {
				t2 += t0
			}
		}
		t0, t1 = t1, t2
	}
	return t2
}

// @lc code=start

// @lc code=end
func Solve91() {
	s := "226"
	fmt.Println(numDecodings1(s))
}
