/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */
package problems

import "fmt"

/*
直接遍历
超时
*/
func rangeBitwiseAnd1(m int, n int) int {
	for i := m + 1; i <= n; i++ {
		m = m & i
		fmt.Println(m)
	}
	return m
}

/*
0&任何数 == 0
*/
func rangeBitwiseAnd2(m int, n int) int {
	for i := m + 1; i <= n; i++ {
		m = m & i
		if m == 0 {
			break
		}
	}
	return m
}

/*
如果存在n使m<2^n<=n，则结果必为0
*/
func rangeBitwiseAnd3(m int, n int) int {
	i := 1
	for i <= m {
		i <<= 1
	}
	if i <= n {
		return 0
	}
	for i := m + 1; i <= n; i++ {
		m = m & i
	}
	return m
}

/*
不断将n的最低位1置为0
*/
func rangeBitwiseAnd4(m int, n int) int {
	for n > m {
		n = n & (n - 1)
	}
	return n
}

/*
同时右移直到相等（右移表示该位有0有1结果必为0）
*/
func rangeBitwiseAnd5(m int, n int) int {
	var zeros int
	for m != n {
		m >>= 1
		n >>= 1
		zeros++
	}
	return m << zeros
}

// @lc code=start

// @lc code=end

func Solve201() {
	m, n := 9, 16
	fmt.Println(rangeBitwiseAnd1(m, n))
}
