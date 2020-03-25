package interview

import (
	"fmt"
)

// map两次遍历
func firstUniqChar1(s string) byte {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	for i, v := range s {
		if m[v] == 1 {
			return s[i]
		}
	}
	return ' '
}

// 如果数据为字母，范围固定，可以用数组
func firstUniqChar2(s string) byte {
	var m [26]int
	for _, v := range s {
		m[v-'a']++
	}
	for i, v := range s {
		if m[v-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func SolveOffer50() {
	s := "abaccdeff"
	fmt.Println(firstUniqChar1(s))
}
