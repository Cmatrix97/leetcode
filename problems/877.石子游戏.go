/*
 * @lc app=leetcode.cn id=877 lang=golang
 *
 * [877] 石子游戏
 */
package problems

import "fmt"

/*
动态规划
博弈问题通用解法
dp[i][j][0/1]
i,j表示从第i堆开始，第j堆结束
0表示先手能获得的最高分
1表示后手能获得的最高分
*/
func stoneGame1(piles []int) bool {
	N := len(piles)
	dp := make([][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < N; i++ {
		dp[i][i][0] = piles[i]
		dp[i][i][1] = 0
	}
	for l := 2; l <= N; l++ {
		for i := 0; i <= N-l; i++ {
			j := l + i - 1
			left := piles[i] + dp[i+1][j][1]
			right := piles[j] + dp[i][j-1][1]
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	return dp[0][N-1][0] > dp[0][N-1][1]
}

func stoneGame2(piles []int) bool {
	return true
}

// @lc code=start

// @lc code=end
func Solve877() {
	piles := []int{5, 3, 4, 5}
	fmt.Println(stoneGame1(piles))
}
