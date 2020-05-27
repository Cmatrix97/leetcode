/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */
package problems

import "fmt"

// 模拟 有10先用10 贪心
func lemonadeChange1(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=start
// @lc code=end

func Solve860() {
	bills := []int{5, 5, 5, 10, 20}
	fmt.Println(lemonadeChange1(bills))
}
