/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
package problems

import "fmt"

func generateParenthesis1(n int) []string {
	var res []string
	var dfs func(left, right int, temp []byte)
	dfs = func(left, right int, temp []byte) {
		if len(temp) == n*2 {
			if left == right {
				res = append(res, string(temp))
			}
			return
		}

		dfs(left+1, right, append(temp, '('))
		if left != right {
			dfs(left, right+1, append(temp, ')'))
		}

	}
	dfs(0, 0, []byte{})
	return res
}

// @lc code=start
// @lc code=end

func Solve22() {
	fmt.Println(generateParenthesis1(3))
}
