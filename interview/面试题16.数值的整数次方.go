package interview

import (
	"fmt"
	"math"
)

// 超时291/304
func myPow1(x float64, n int) float64 {
	var res float64 = 1
	abs := math.Abs(float64(n))
	for abs != 0 {
		res *= x
		abs--
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

// 递归，二分思想
func myPow2(x float64, n int) float64 {
	if n == -1 {
		return 1 / x
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	half := myPow2(x, n>>1)
	mod := myPow2(x, n&1)
	return half * half * mod
}

// 快速幂，将n转为2进制，每次右移一位，x每次平方，如果n当前位为1则res*=x
// Go中负数右移最大为-1
func myPow3(x float64, n int) float64 {
	var res float64 = 1
	t := int(math.Abs(float64(n)))
	for t != 0 {
		if t&1 == 1 {
			res *= x
		}
		x *= x
		t >>= 1
	}
	if n < 0 {
		return 1 / res
	}
	return res
}

func SolveOffer16() {
	x := 2.00000
	n := -2
	fmt.Println(myPow3(x, n))
}
