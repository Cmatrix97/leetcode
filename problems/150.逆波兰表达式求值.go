/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */
package problems

import (
	"fmt"
	"strconv"
)

/*
栈
*/
func evalRPN1(tokens []string) int {
	var stack []int
	isNum := func(s string) bool {
		if s == "+" || s == "-" || s == "*" || s == "/" {
			return false
		}
		return true
	}
	for _, v := range tokens {
		if isNum(v) {
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		} else {
			left, right := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var ans int
			switch v {
			case "+":
				ans = left + right
			case "-":
				ans = left - right
			case "*":
				ans = left * right
			case "/":
				ans = left / right
			}
			stack = append(stack, ans)
		}
	}
	return stack[0]
}

// @lc code=start

// @lc code=end

func Solve150() {
	tokens := []string{
		"4", "13", "5", "/", "+",
	}
	fmt.Println(evalRPN1(tokens))
}
