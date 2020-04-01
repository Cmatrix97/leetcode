package interview

import (
	"fmt"
	"math"
)

func maxProfit1(prices []int) int {
	min, max := math.MaxInt64, 0
	for _, v := range prices {
		if v-min > max {
			max = v - min
		}
		if v < min {
			min = v
		}
	}
	return max
}

func SolveOffer63() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(prices))
}
