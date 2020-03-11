package interview

import "fmt"

func cuttingRope1_2(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans, mod := 1, 1000000007
	for n > 4 {
		ans %= mod
		ans *= 3
		n -= 3
	}
	return n * ans % mod
}

func SolveOffer14_2() {
	n := 120
	fmt.Println(cuttingRope1_2(n))
}
