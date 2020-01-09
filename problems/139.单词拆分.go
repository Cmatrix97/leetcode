/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */
package problems

import (
	"fmt"
	"strings"
)

/*
递归超时
*/
func wordBreak1(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range wordDict {
		if strings.HasPrefix(s, v) {
			if wordBreak1(s[len(v):], wordDict) {
				return true
			}
		}
	}
	return false
}

/*
动态规划
dp[i]表示s中前i个字符构成的子串是否可拆
*/
func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, v := range wordDict {
			if i >= len(v) && s[i-len(v):i] == v && dp[i-len(v)] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

// @lc code=start

// @lc code=end

func Solve139() {
	s := "abcd"
	wordDict := []string{
		"a",
		"abc",
		"b",
		"cd",
	}
	fmt.Println(wordBreak1(s, wordDict))
}
