/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */
package problems

import "fmt"

/*
动态规划
dp[i][j]表示是s[i:j]的最长回文子序列长度
(1)s[i] == s[j]
dp[i][j] = dp[i+1][j-1]+2
(2)s[i] != s[j]
dp[i][j] = max{dp[i+1][j], dp[i][j-1]}
初始
dp[i][i] = 1
结果
dp[0][len(s)-1]
*/
func longestPalindromeSubseq1(s string) int {
	N := len(s)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		dp[i][i] = 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := N - 1; i >= 0; i-- {
		for j := i + 1; j < N; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][N-1]
}

/*
降维
dp[i]表示前i个元素的最长回文子序列长度
*/
func longestPalindromeSubseq2(s string) int {
	N := len(s)
	dp := make([]int, N)
	f := make([]int, 2)
	for i, _ := range dp {
		dp[i] = 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := N; i >= 0; i-- {
		f[0] = 0
		for j := i + 1; j < N; j++ {
			f[1] = dp[j]
			if s[i] == s[j] {
				dp[j] = f[0] + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			f[0] = f[1]
		}
	}
	return dp[N-1]
}

// @lc code=start

// @lc code=end
func Solve516() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq1(s))
}
