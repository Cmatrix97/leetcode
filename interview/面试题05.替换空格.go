package interview

import (
	"bytes"
	"fmt"
	"strings"
)

// silce
func replaceSpace1(s string) string {
	var res = []byte{}
	for _, v := range s {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, byte(v))
		}
	}
	return string(res)
}

// bytes.Buffer
func replaceSpace2(s string) string {
	var buf bytes.Buffer
	for _, v := range s {
		if v == ' ' {
			buf.WriteString("%20")
			continue
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

// strings库函数
func replaceSpace3(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

// S(n) = O(1)的思路
// 由于Go无法原地修改string，空间复杂度降不到O(1)
func replaceSpace4(s string) string {
	var count int
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	str := make([]byte, len(s)+2*count)
	for i, j := len(str)-1, len(s)-1; j >= 0;j-- {
		if s[j] == ' ' {
			str[i-2], str[i-1], str[i] = '%', '2', '0'
			i -= 3
		} else {
			str[i] = s[j]
			i -= 1
		}
	}
	return string(str)
}

func SolveOffer05() {
	s := "We are happy."
	fmt.Println(replaceSpace1(s))
}
