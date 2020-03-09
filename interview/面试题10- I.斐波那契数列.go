package interview

import "fmt"

func fib1(n int) int {
	n0, n1 := 0, 1
	for i := 0; i < n; i++ {
		n0, n1 = n1, (n0+n1)%1000000007
	}
	return n0
}

func SolveOffer10_1() {
	n := 100
	fmt.Println(fib1(n))
}
