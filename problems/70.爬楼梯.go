/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
package problems

import (
	"fmt"
	"math"
)

/*
斐波那契数列
*/
func climbStairs1(n int) int {
	w0, w1 := 1, 1
	for i := 2; i <= n; i++ {
		w0, w1 = w1, w0+w1
	}
	return w1
}

/*
公式法
Fn = 1/sqrt5 * {[(1+sqrt5)/2]^(n+1) - [(1-sqrt5)/2]^(n+1)}
*/
func climbStairs2(n int) int {
	sqrt5 := math.Sqrt(5)
	pow := func(x float64, y int) float64 {
		var res float64 = 1
		for i := 0; i < y; i++ {
			res *= x
		}
		return res
	}
	fibn := pow((1+sqrt5)/2, n+1) - pow((1-sqrt5)/2, n+1)
	return int(fibn / sqrt5)
}

// @lc code=start

// @lc code=end
func Solve70() {
	n := 9
	fmt.Println(climbStairs1(n))
}
