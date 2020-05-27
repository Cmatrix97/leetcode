/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */
package problems

import (
	"fmt"
)

// 回溯
func solveNQueens1(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	isValid := func(board [][]byte, row, col int) bool {
		n := len(board)
		for i := 0; i < n; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = 'Q'
			backtrack(row + 1)
			board[row][col] = '.'
		}
	}
	backtrack(0)
	return res
}

// 数组记录
func solveNQueens2(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	na := make([]bool, 2*n-1)
	pie := make([]bool, 2*n-1)
	shu := make([]bool, n)
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}
		for col := 0; col < n; col++ {
			i, j := row+col, col-row+n-1
			if pie[i] || na[j] || shu[col] {
				continue
			}
			board[row][col], pie[i], na[j], shu[col] = 'Q', true, true, true
			backtrack(row + 1)
			board[row][col], pie[i], na[j], shu[col] = '.', false, false, false
		}
	}
	backtrack(0)
	return res
}

// 位运算记录
func solveNQueens3(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	shu, pie, na := 0, 0, 0
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}
		for col := 0; col < n; col++ {
			i, j := row+col, col-row+n-1
			if (pie>>i|na>>j|shu>>col)&1 != 0 {
				continue
			}
			board[row][col], shu, pie, na = 'Q', shu|(1<<col), pie|(1<<i), na|(1<<j)
			backtrack(row + 1)
			board[row][col], shu, pie, na = '.', shu&^(1<<col), pie&^(1<<i), na&^(1<<j)
		}
	}
	backtrack(0)
	return res
}

func solveNQueens4(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	shu, pie, na := 0, 0, 0
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}
		available := ((1 << n) - 1) & ^(shu | (pie >> row) | (na >> (n - 1 - row)))
		for available != 0 {
			p := available & -available
			available = available &^ p
			temp, col := p, 0
			for temp != 1 {
				temp, col = temp>>1, col+1
			}
			board[row][col], shu, pie, na = 'Q', shu|p, pie|(p<<row), na|(p<<(n-1-row))
			backtrack(row + 1)
			board[row][col], shu, pie, na = '.', shu&^p, pie&^(p<<row), na&^(p<<(n-1-row))
		}
	}
	backtrack(0)
	return res
}

func solveNQueens5(n int) [][]string {
	var res [][]string
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	var backtrack func(row, shu, pie, na int)
	backtrack = func(row, shu, pie, na int) {
		if row == n {
			var temp []string
			for _, v := range board {
				temp = append(temp, string(v))
			}
			res = append(res, temp)
			return
		}
		available := ((1 << n) - 1) & ^(shu | pie | na)
		for available != 0 {
			p := available & -available
			available = available &^ p
			temp, col := p, 0
			for temp != 1 {
				temp, col = temp>>1, col+1
			}
			board[row][col] = 'Q'
			backtrack(row+1, shu^p, (pie^p)>>1, (na^p)<<1)
			board[row][col] = '.'
		}
	}
	backtrack(0, 0, 0, 0)
	return res
}

// @lc code=start

// @lc code=end

func Solve51() {
	n := 8
	fmt.Println(solveNQueens1(n))
}
