/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */
package problems

import (
	"fmt"
	"math"
)

/*
中心扩展算法
遍历两次，第一次为奇回文，第二次为偶回文
*/
func longestPalindrome1(s string) string {
	var maxLength int
	var maxString string
	for i := range s {
		count := 1
		p, q := i-1, i+1
		for p >= 0 && q < len(s) {
			if s[p] != s[q] {
				break
			}
			count += 2
			p--
			q++
		}
		if count > maxLength {
			maxLength = count
			maxString = s[p+1 : q]
		}
	}
	for post := 1; post < len(s); post++ {
		p, q := post-1, post
		count := 0
		for p >= 0 && q < len(s) {
			if s[p] != s[q] {
				break
			}
			count += 2
			p--
			q++
		}
		if count > maxLength {
			maxLength = count
			maxString = s[p+1 : q]
		}
	}
	return maxString
}

/*
第一种方法的优化，避免了重复遍历及字符串赋值
*/
func longestPalindrome2(s string) string {
	if len(s) < 1 {
		return ""
	}
	var start, end int
	for i := range s {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end-start+1 {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

/*
动态规划法
如果s[i:j]为回文串，p(i,j) = true，否则为false
p(i,j) = p(i+1,j-1) && (s[i] == s[j)
*/
func longestPalindrome3(s string) string {
	var res string
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for start := n - 1; start >= 0; start-- {
		for end := start; end < n; end++ {
			dp[start][end] = s[start] == s[end] && (end-start < 2 || dp[start+1][end-1])
			if dp[start][end] && end-start+1 > len(res) {
				res = s[start : end+1]
			}
		}
	}
	return res
}

/*
Manacher算法
时间复杂度O(n)
*/
func preProcess(s string) string {
	res := "^"
	if len(s) == 0 {
		return res + "$"
	}
	for _, v := range s {
		res += "#" + string(v)
	}
	res += "#$"
	return res
}

func longestPalindrome4(s string) string {
	T := preProcess(s)
	length := len(T)
	P := make([]int, length)
	var C, R int
	for i := 1; i < length-1; i++ {
		i_mirror := 2*C - i
		if R > i {
			P[i] = int(math.Min(float64(R-i), float64(P[i_mirror])))
		} else {
			P[i] = 0
		}
		for T[i+1+P[i]] == T[i-1-P[i]] {
			P[i]++
		}
		if i+P[i] > R {
			C, R = i, i+P[i]
		}
	}
	var maxLen, centerIndex int
	for i := 1; i < length-1; i++ {
		if P[i] > maxLen {
			maxLen = P[i]
			centerIndex = i
		}
	}
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

// @lc code=start

// @lc code=end

func Solve5() {
	s := "babad"
	fmt.Println(longestPalindrome1(s))
}
