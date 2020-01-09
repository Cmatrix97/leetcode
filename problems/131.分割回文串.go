/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */
package problems

import "fmt"

/*
回溯
*/
func partition1_131(s string) [][]string {
	var res [][]string
	var temp []string
	isPalindromic := func(s string, left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			t := make([]string, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := start; i < len(s); i++ {
			if !isPalindromic(s, start, i) {
				continue
			}
			temp = append(temp, s[start:i+1])
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0)
	return res
}

/*
动态规划判断回文串
第5题方法3
*/
func partition2_131(s string) [][]string {
	length := len(s)
	var res [][]string
	var temp []string
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for start := length - 1; start >= 0; start-- {
		for end := start; end < length; end++ {
			dp[start][end] = s[start] == s[end] && (end-start < 2 || dp[start+1][end-1])
		}
	}
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == length {
			t := make([]string, len(temp))
			copy(t, temp)
			res = append(res, t)
			return
		}
		for i := start; i < len(s); i++ {
			if !dp[start][i] {
				continue
			}
			temp = append(temp, s[start:i+1])
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0)
	return res
}

// @lc code=start

// @lc code=end

func Solve131() {
	s := "aab"
	res := partition1_131(s)
	fmt.Println(res)
}
