/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package problems

import "fmt"

// dfs
func letterCombinations1(digits string) []string {
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var res []string
	var dfs func(index int, tmp []byte)
	dfs = func(index int, tmp []byte) {
		if index == len(digits) {
			res = append(res, string(tmp))
			return
		}

		for _, v := range []byte(m[digits[index]]) {
			dfs(index+1, append(tmp, v))
		}
	}
	if len(digits) != 0 {
		dfs(0, []byte{})
	}
	return res
}

// @lc code=start
// @lc code=end

func Solve17() {
	digits := "23"
	fmt.Println(letterCombinations1(digits))
}
