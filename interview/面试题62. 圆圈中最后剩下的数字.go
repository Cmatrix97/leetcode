package interview

import (
	"fmt"
)

// 暴力法 超时
func lastRemaining1(n int, m int) int {
	queue := make([]int, n)
	for i := 0; i < n; i++ {
		queue[i] = i
	}
	for len(queue) > 1 {
		for i := 0; i < m-1; i++ {
			el := queue[0]
			queue = queue[1:]
			queue = append(queue, el)
		}
		queue = queue[1:]
	}
	return queue[0]
}

func lastRemaining2(n, m int) int {
	var last int
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

func SolveOffer62() {
	n, m := 5, 3
	fmt.Println(lastRemaining1(n, m))
}
