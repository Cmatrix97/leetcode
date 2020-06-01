package interview

import (
	"fmt"
	"math"
)

func strToInt1(str string) int {
	var i int
	for i = 0; i < len(str) && str[i] == ' '; i++ {
	}
	if i == len(str) {
		return 0
	}
	flag := true
	if str[i] == '-' {
		flag = false
	}
	if str[i] == '-' || str[i] == '+' {
		i++
	}
	var next int
	var result int
	for i < len(str) && str[i] >= '0' && str[i] <= '9' {
		next = int(str[i] - '0')
		if result > math.MaxInt32/10 || result == math.MaxInt32/10 && next > 7 {
			if flag {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = result*10 + next
		i++
	}
	if !flag {
		return -result
	}
	return result
}

func SolveOffer67() {
	str := "  -42"
	fmt.Println(strToInt1(str))
}
