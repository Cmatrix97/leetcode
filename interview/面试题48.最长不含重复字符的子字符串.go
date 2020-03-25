package interview

import "fmt"

// 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	var ans int
	start := -1
	m := make(map[rune]int)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for end, v := range s {
		if idx, ok := m[v]; ok {
			start = max(start, idx)
		}
		ans = max(ans, end-start)
		m[v] = end
	}
	return ans
}

func SolveOffer48() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring1(s))
}
