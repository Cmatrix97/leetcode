/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

/*栈的简单应用
1.如果当前字符为左括号时，则将其压入栈中
2.如果遇到右括号时：
(1)如栈不为空且为对应的左半边括号，则取出栈顶元素，继续循环
(2)若此时栈为空，则直接返回false
(3)若不为对应的左半边括号，反之返回false
*/
package problems

import "fmt"

func isValid1(s string) bool {
	stack, top := []rune{}, 0
	stack = make([]rune, 10000)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack[top] = v
			top++
		} else if v == ')' || v == ']' || v == '}' {
			if top == 0 {
				return false
			}
			top--
			temp := stack[top]
			if (v == ')' && temp != '(') || (v == ']' && temp != '[') || (v == '}' && temp != '{') {
				return false
			}
		}
	}
	if top != 0 {
		return false
	}
	return true
}

// @lc code=start
func isValid(s string) bool {
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []rune{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		}
		if v == ')' || v == ']' || v == '}' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if t := m[v]; t != top {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end

func Solution20() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
