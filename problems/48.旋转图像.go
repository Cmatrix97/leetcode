/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */
package problems

import "fmt"

/*
先转置，再逐行倒置
*/
func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

/*
切分成四部分，再逐个交换
12 --> 31
34 --> 42
行 i ->   j   -> n-1-i -> n-1-j -> i
列 j -> n-1-i -> n-1-j ->   i   -> j
*/
func rotate2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}

// @lc code=start

// @lc code=end

func Solve48() {
	matrix := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate1(matrix)
	fmt.Println(matrix)
}
