/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
package problems

import (
	"fmt"
	"strings"
)

/*二次遍历（算ToLower三次）
 */
func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	arr := []rune{}
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			arr = append(arr, c)
		}
	}
	if len(arr) == 0 {
		return true
	}
	for i := 0; i <= len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

/*双指针一次遍历
空间复杂度O(1)
*/
func isPalindrome2(s string) bool {
	isalnum := func(c byte) bool {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			return true
		}
		return false
	}
	i, j := 0, len(s)-1
	for i < j {
		if !isalnum(s[i]) {
			i++
			continue
		}
		if !isalnum(s[j]) {
			j--
			continue
		}
		if (s[i] | 0x20) != (s[j] | 0x20) {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=start

// @lc code=end

func Solve125() {
	// s := "A man, a plan, a canal: Panama"
	// s := "race a car"
	s := " "
	fmt.Println(isPalindrome1((s)))
}
