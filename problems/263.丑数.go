/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */
package problems

import "fmt"

func isUgly(num int) bool {
	for num != 1 {
		temp := num
		if num%2 == 0 {
			num /= 2
		}
		if num%3 == 0 {
			num /= 3
		}
		if num%5 == 0 {
			num /= 5
		}
		if num == temp {
			break
		}
	}
	if num != 1 {
		return false
	}
	return true
}

// @lc code=start

// @lc code=end

func Solve263() {
	num := 14
	fmt.Println(isUgly(num))
}
