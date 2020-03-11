package interview

import (
	"fmt"
	"math"
)

// 暴力，只能解决小数据
func printNumbers1(n int) []int {
	max := int(math.Pow10(n))
	res := make([]int, 0, max-1)
	for i := 1; i < max; i++ {
		res = append(res, i)
	}
	return res
}

// 全排列 cacl工具函数将[]int转为要求的int
func printNumbers2(n int) []int {
	max := int(math.Pow10(n))
	res := make([]int, 0, max)
	arr := make([]int, n)
	cacl := func(arr []int) int {
		var sum int
		for i := len(arr) - 1; i >= 0; i-- {
			sum += arr[i] * int(math.Pow10(len(arr)-1-i))
		}
		return sum
	}
	var recurse func(idx int)
	recurse = func(idx int) {
		if idx == n {
			tmp := cacl(arr)
			res = append(res, tmp)
			return
		}
		for i := 0; i < 10; i++ {
			arr[idx] = i
			recurse(idx + 1)
		}
	}
	recurse(0)
	return res[1:]
}

func SolveOffer17() {
	n := 5
	fmt.Println(printNumbers2(n))
}
