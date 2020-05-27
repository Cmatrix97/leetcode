/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
package problems

import "fmt"

// 二分法
func mySqrt1(x int) int {
	lo, hi := 0, x
	ans := -1
	for lo <= hi {
		mi := lo + (hi-lo)>>1
		if mi*mi > x {
			hi = mi - 1
		} else {
			ans = mi
			lo = mi + 1
		}
	}
	return ans
}

// 牛顿迭代法
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	}

	c := float64(x)
	xi := float64(x)
	for {
		xi = (xi + c/xi) / 2
		if xi*xi-c < 10e-1 {
			break
		}
	}
	return int(xi)
}

// @lc code=start

// @lc code=end

func Solve69() {
	x := 8
	fmt.Println(mySqrt1(x))
}
