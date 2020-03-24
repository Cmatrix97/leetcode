package interview

import (
	"fmt"
	"strconv"
)

func translateNum1(num int) int {
	str := strconv.Itoa(num)
	var recurse func(str string) int
	recurse = func(str string) int {
		if len(str) <= 1 {
			return 1
		}
		if str[0] == '1' || (str[0] == '2' && str[1] <= '5') {
			return recurse(str[1:]) + recurse(str[2:])
		} else {
			return recurse(str[1:])
		}
	}
	return recurse(str)
}

func SolveOffer46() {
	num := 12258
	fmt.Println(translateNum1(num))
}
