/*
 * @lc app=leetcode.cn id=688 lang=golang
 *
 * [688] “马”在棋盘上的概率
 */
package problems

import "fmt"

// 递归 超时
func knightProbability1(N int, K int, r int, c int) float64 {
	if r < 0 || r >= N || c < 0 || c >= N {
		return 0
	}
	if K == 0 {
		return 1
	}
	var total float64
	total += knightProbability1(N, K-1, r+2, c+1) + knightProbability1(N, K-1, r+2, c-1)
	total += knightProbability1(N, K-1, r-2, c+1) + knightProbability1(N, K-1, r-2, c-1)
	total += knightProbability1(N, K-1, r+1, c+2) + knightProbability1(N, K-1, r+1, c-2)
	total += knightProbability1(N, K-1, r-1, c+2) + knightProbability1(N, K-1, r-1, c-2)
	return total / 8
}

// 三维dp
func knightProbability2(N int, K int, r int, c int) float64 {
	moves := [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	dp := make([][][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]float64, K+1)
		}
	}
	for k := 0; k <= K; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if k == 0 {
					dp[i][j][k] = 1
				} else {
					for _, move := range moves {
						nextX, nextY := i+move[0], j+move[1]
						if nextX < 0 || nextX >= N || nextY < 0 || nextY >= N {
							continue
						}
						dp[i][j][k] += dp[nextX][nextY][k-1]
					}
					dp[i][j][k] /= 8
				}
			}
		}
	}
	return dp[r][c][K]
}

// @lc code=start
// 空间复杂度降为O(n^2)
func knightProbability(N int, K int, r int, c int) float64 {
	moves := [][]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
	temp, dp := make([][]float64, N), make([][]float64, N)
	for i := 0; i < N; i++ {
		temp[i], dp[i] = make([]float64, N), make([]float64, N)
		for j := 0; j < N; j++ {
			dp[i][j] = 1
		}
	}
	for m := 1; m <= K; m++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for _, move := range moves {
					nextX, nextY := i+move[0], j+move[1]
					if nextX < 0 || nextX >= N || nextY < 0 || nextY >= N {
						continue
					}
					temp[i][j] += dp[nextX][nextY]
				}
				temp[i][j] /= 8
			}
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				dp[i][j], temp[i][j] = temp[i][j], 0
			}
		}
	}
	return dp[r][c]
}

// @lc code=end

func Solve688() {
	N, K, r, c := 3, 2, 0, 0
	fmt.Println(knightProbability(N, K, r, c))
}
