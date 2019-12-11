/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */
package problems

import "fmt"

/*
相加 -> a ^ b
进位 -> (a & b) << 1
将进位继续相加直到进位为0
*/
func getSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// @lc code=start

// @lc code=end

func Solve371() {
	a, b := -2, 3
	fmt.Println(getSum(a, b))
}
