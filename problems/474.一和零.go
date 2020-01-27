/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */
package problems

import "fmt"

/*
二维费用的背包问题
dp[i][j]表示使用i个0和j个1最多能拼出的字符串数
*/
func findMaxForm1(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	computeZeroOne := func(str string) (count0, count1 int) {
		for _, v := range str {
			if v == '0' {
				count0++
			} else {
				count1++
			}
		}
		return
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for k := 0; k < len(strs); k++ {
		count0, count1 := computeZeroOne(strs[k])
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
	}
	return dp[m][n]
}

// @lc code=start

// @lc code=end
func Solve474() {
	array := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm1(array, m, n))
}
