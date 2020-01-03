/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */
package problems

import "fmt"

/*
回溯DFS
*/
func exist1(board [][]byte, word string) bool {
	M, N := len(board), len(board[0])
	if len(word) > M*N {
		return false
	}
	var used [][]bool
	used = make([][]bool, M)
	for i := 0; i < M; i++ {
		used[i] = make([]bool, N)
	}
	direction := [][]int{
		{-1, 0},
		{0, -1},
		{0, 1},
		{1, 0},
	}
	isValid := func(row, col int, char byte) bool {
		if row < 0 || row >= M || col < 0 || col >= N {
			return false
		}
		if used[row][col] {
			return false
		}
		return true
	}
	var backtrack func(row, col, index int) bool
	backtrack = func(row, col, index int) bool {
		if index == len(word)-1 {
			return board[row][col] == word[index]
		}
		if board[row][col] == word[index] {
			used[row][col] = true
			for i := 0; i < 4; i++ {
				nextRow, nextCol := row+direction[i][0], col+direction[i][1]
				if !isValid(nextRow, nextCol, word[index+1]) {
					continue
				}
				if backtrack(nextRow, nextCol, index+1) {
					return true
				}
			}
			used[row][col] = false
		}
		return false
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=start

// @lc code=end

func Solve79() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist1(board, word))
}
