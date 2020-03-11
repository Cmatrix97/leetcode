package interview

import (
	"fmt"
)

// dp
// dp[i]存储长度为i的绳子最大切割乘积（且至少剪一次）
// 每次比较max{dp[j], j}*(i-j) j:1->i-1
// 由于已经至少切了一刀，因此应比较dp[j]（至少剪一次）和j（不剪）
func cuttingRope1(n int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], (i-j)*max(dp[j], j))
		}
	}
	return dp[n]
}

// 贪心，3为最优
func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		ans %= 1000000007
		ans *= 3
		n /= 3
	}
	return n * ans
}

func SolveOffer14_1() {
	n := 120
	fmt.Println(cuttingRope2(n))
}
