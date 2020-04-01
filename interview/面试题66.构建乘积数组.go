package interview

import "fmt"

// O(n^2) 超时
func constructArr1(a []int) []int {
	res := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = 1
		for j := 0; j < len(a); j++ {
			if i != j {
				res[i] *= a[j]
			}
		}
	}
	return res
}

// 左右子数组
func constructArr2(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	left, right := make([]int, len(a)), make([]int, len(a))
	res := make([]int, len(a))
	left[0], right[len(right)-1] = 1, 1
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1] * a[i-1]
		right[len(a)-1-i] = right[len(a)-i] * a[len(a)-i]
	}
	for i := 0; i < len(a); i++ {
		res[i] = left[i] * right[i]
	}
	return res
}

func SolveOffer66() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr1(a))
}
