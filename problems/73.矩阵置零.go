/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */
package problems

import "fmt"

/*
每行或每列第一位用来表示该行/列是否置0
matrix[0][0]不能同时表示第一行和第一列
需要额外设置一个bool变量
*/
func setZeroes1(matrix [][]int) {
	M, N := len(matrix), len(matrix[0])
	var isCol0 bool
	for i := 0; i < M; i++ {
		if matrix[i][0] == 0 {
			isCol0 = true
		}
		for j := 1; j < N; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := 1; i < M; i++ {
		for j := 1; j < N; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0; j < N; j++ {
			matrix[0][j] = 0
		}
	}
	if isCol0 {
		for i := 0; i < M; i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=start

// @lc code=end

func Solve73() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes1(matrix)
	fmt.Println(matrix)
}
