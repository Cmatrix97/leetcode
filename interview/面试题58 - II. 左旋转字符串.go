package interview

import (
	"fmt"
	"unsafe"
)

// 三次reverse
func reverseLeftWords1(s string, n int) string {
	reverse := func(bytes []byte) {
		for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
	}
	strs := []byte(s)
	reverse(strs[:n])
	reverse(strs[n:])
	reverse(strs)
	return *(*string)(unsafe.Pointer(&strs))
}

func SolveOffer58_2() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords1(s, k))
}
