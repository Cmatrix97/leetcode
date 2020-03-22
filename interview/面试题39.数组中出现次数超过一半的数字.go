package interview

import (
	"fmt"
	"math"
)

// 投票算法
func majorityElement1(nums []int) int {
	candidate, votes := math.MinInt64, 0
	for _, num := range nums {
		if votes == 0 {
			candidate = num
		}
		if num == candidate {
			votes++
		} else {
			votes--
		}
	}
	return candidate
}

func SolveOffer39() {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement1(nums))
}
