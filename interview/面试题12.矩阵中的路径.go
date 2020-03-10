package interview

import "fmt"

// dfs
// 回溯下一位置时将当前位置置为0（不可能出现的值）表示已经走过
func exist1(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	M, N := len(board), len(board[0])
	var dfs func(row, col int, idx int) bool
	dfs = func(row, col int, idx int) bool {
		if board[row][col] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		tmp := board[row][col]
		board[row][col] = '0'
		if (row-1 >= 0 && dfs(row-1, col, idx+1)) || (row+1 < M && dfs(row+1, col, idx+1)) || (col-1 >= 0 && dfs(row, col-1, idx+1)) || (col+1 < N && dfs(row, col+1, idx+1)) {
			return true
		}
		board[row][col] = tmp
		return false
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func SolveOffer12() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist1(board, word))
}
