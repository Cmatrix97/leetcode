package interview

import "fmt"

// dfs
// 外层循环缩小范围，每圈行数和列数-2，当其一小于等于0时结束
// 每圈的起始坐标为matrix[i][i]，每圈i+1
// 如果只有一行或一列，直接在中途跳出递归（因为不会向反方向走）
func spiralOrder1(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	rows, cols := len(matrix), len(matrix[0])
	start := 0
	var print func(x, y int)
	print = func(x, y int) {
		res = append(res, matrix[x][y])
		if x == start && y < start+cols-1 {
			print(x, y+1)
		}
		if x < start+rows-1 && y == start+cols-1 {
			print(x+1, y)
		}
		if rows == 1 || cols == 1 {
			return
		}
		if x == start+rows-1 && y > start {
			print(x, y-1)
		}
		if x > start+1 && y == start {
			print(x-1, y)
		}
	}
	for rows > 0 && cols > 0 {
		print(start, start)
		rows, cols, start = rows-2, cols-2, start+1
	}
	return res
}

func SolveOffer29() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	res := spiralOrder1(matrix)
	fmt.Println(res)
}
