package interview

import (
	"fmt"
	"math"
)

func sumNums1(n int) int {
	return int(math.Pow(float64(n), 2)+float64(n)) >> 1
}

func SolveOffer64() {
	n := 10
	fmt.Println(sumNums1(n))
}
