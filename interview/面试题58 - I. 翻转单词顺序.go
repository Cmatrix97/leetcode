package interview

import (
	"bytes"
	"fmt"
	"strings"
)

// 两次reverse
func reverseWords1(s string) string {
	if len(s) == 0 {
		return s
	}
	reverse := func(bytes []byte) {
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
	}
	strs := []byte(s)
	reverse(strs)
	var start, end int
	var buf bytes.Buffer
	for start != len(s) {
		if strs[start] == ' ' {
			start++
			end++
		} else if end == len(s) || strs[end] == ' ' {
			reverse(strs[start:end])
			buf.Write(strs[start:end])
			buf.Write([]byte{' '})
			start, end = end, end+1
		} else {
			end++
		}
	}
	return strings.TrimSpace(buf.String())
}

// 库函数
func reverseWords2(s string) string {
	parts := strings.Split(s, " ")
	size := len(parts)
	var rparts []string
	for i := size - 1; i >= 0; i-- {
		if parts[i] != "" {
			rparts = append(rparts, parts[i])
		}
	}
	return strings.Join(rparts, " ")
}

func SolveOffer58_1() {
	s := "the sky is   blue "
	fmt.Println(reverseWords1(s))
}
