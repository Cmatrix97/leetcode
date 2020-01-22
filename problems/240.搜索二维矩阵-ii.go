/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package problems

import (
	"fmt"
)

/*
暴力法
*/
func searchMatrix1(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

/*
从左下角或右上角开始
*/
func searchMatrix2(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0
	for row >= 0 && col <= len(matrix[0])-1 {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

// @lc code=start

// @lc code=end
func Solve240() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 20
	fmt.Println(searchMatrix1(matrix, target))
}
