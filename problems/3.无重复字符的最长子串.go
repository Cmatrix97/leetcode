/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
package problems

import (
	"fmt"
	"math"
)

/*
通过hashmap记录元素的下标
遍历字符串
1.(1)当前元素已存在，count=当前下标-上次出现的下标，同时清空所有下标<=上次出现的下标的元素
  (2)当前元素不存在，count+1
2.更新元素下标，比较当前count和最大值
*/
func lengthOfLongestSubstring1(s string) int {
	m := make(map[rune]int)
	count, max := 0, 0
	for curIndex, v := range s {
		if preIndex, ok := m[v]; ok {
			count = curIndex - preIndex
			for del, val := range m {
				if val <= preIndex {
					delete(m, del)
				}
			}
		} else {
			count++
		}
		m[v] = curIndex
		if count > max {
			max = count
		}
	}
	return max
}

/*
hashset实现滑动窗口
(1)若新元素不存在于set中时判断是否更新ans，并将j右移
(2)否则将i右移
*/
func lengthOfLongestSubstring2(s string) int {
	ans, i, j, n := 0, 0, 0, len(s)
	m := make(map[byte]struct{})
	for i < n && j < n {
		if _, ok := m[s[j]]; !ok {
			m[s[j]] = struct{}{}
			j++
			if j-i > ans {
				ans = j - i
			}
		} else {
			delete(m, s[i])
			i++
		}
	}
	return ans
}

/*
hashmap实现滑动窗口
基于第一种方法的优化，通过start指针避免了清空map的操作
*/
func lengthOfLongestSubstring3(s string) int {
	var ans, start int
	m := make(map[rune]int)
	for end, v := range s {
		if index, ok := m[v]; ok {
			start = int(math.Max(float64(start), float64(index)))
		}
		ans = int(math.Max(float64(ans), float64(end-start+1)))
		m[v] = end + 1
	}
	return ans
}

// @lc code=start

// @lc code=end

func Solve3() {
	s := "abba"
	fmt.Println(lengthOfLongestSubstring1(s))
}
