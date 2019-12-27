/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */
package problems

import (
	"fmt"
	"math"
)

/*
加法，超时
*/
func divide1(dividend int, divisor int) int {
	isPos := (dividend > 0) == (divisor > 0)
	dividend, divisor = int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	sum, count := 0, 0
	for sum <= dividend {
		sum += divisor
		if (isPos && count == 1<<31-1) || (!isPos && count == 1<<31) {
			return 1<<31 - 1
		}
		count++
	}
	if !isPos {
		return 1 - count
	} else {
		return count - 1
	}
}

/*
二分法
*/
func divide2(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	isPos := (dividend < 0) == (divisor < 0)
	dividend, divisor = int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	var res int
	for dividend >= divisor {
		temp, count := divisor, 1
		for dividend >= temp<<1 {
			temp <<= 1
			count <<= 1
		}
		res += count
		dividend -= temp
	}
	if !isPos {
		res = -res
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve29() {
	dividend, divisor := -2147483648, -1
	fmt.Println(divide1(dividend, divisor))
}
