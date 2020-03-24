package interview

import (
	"fmt"
	"strconv"
)

// 找规律
// 一位数（除了0） 9 * 1 * 1
// 两位数         9 * 10 * 2
// 三位数         9 * 100 * 3
func findNthDigit1(n int) int {
	pow, pos := 1, 1
	for {
		remain := pow * pos * 9
		if n <= remain {
			break
		}
		n, pow, pos = n-remain, pow*10, pos+1
	}
	th, mod := n/pos, n%pos
	if mod == 0 {
		return int(strconv.Itoa(pow + th - 1)[pos-1] - '0')
	}
	return int(strconv.Itoa(pow + th)[mod-1] - '0')
}

func SolveOffer44() {
	n := 65536
	fmt.Println(findNthDigit1(n))
}
