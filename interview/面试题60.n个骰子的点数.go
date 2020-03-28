package interview

import (
	"fmt"
)

// 带备忘录的递归
// f(n, k)表示n个骰子和为k出现的次数
// f(n, k) = f(n-1, k-1) + f(n-1, k-2) + f(n-1, k-3)
//         + f(n-1, k-4) + f(n-1, k-5) + f(n-1, k-6)
// k-1~k-6必须满足大于n-1且小于6*(n-1)
func twoSum1(n int) []float64 {
	total := 1
	for i := 0; i < n; i++ {
		total *= 6
	}
	dict := make(map[[2]int]int)
	var recurse func(n, k int) int
	recurse = func(n, k int) int {
		if n == 1 {
			return 1
		}
		var count int
		for i := k - 6; i < k; i++ {
			if i >= n-1 && i <= 6*(n-1) {
				if v, ok := dict[[2]int{n - 1, i}]; ok {
					count += v
				} else {
					count += recurse(n-1, i)
				}
			}
		}
		dict[[2]int{n, k}] = count
		return count
	}
	res := make([]float64, 5*n+1)
	for i := 0; i < len(res); i++ {
		res[i] = float64(recurse(n, i+n)) / float64(total)
	}
	return res
}

// 一维dp
func twoSum2(n int) []float64 {
	total := 1
	for i := 0; i < n; i++ {
		total *= 6
	}
	res := make([]float64, 6*n+1)
	for i := 1; i <= 6; i++ {
		res[i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			res[j] = 0
			for k := j - 6; k < j; k++ {
				if k >= i-1 && k <= 6*(i-1) {
					res[j] += res[k]
				}
			}
		}
	}
	for i := range res {
		res[i] /= float64(total)
	}
	return res[n:]
}

func SolveOffer60() {
	n := 2
	fmt.Println(twoSum2(n))
}
