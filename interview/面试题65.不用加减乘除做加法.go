package interview

import "fmt"

func add1(a int, b int) int {
	var sum, carry int
	for {
		sum = a ^ b
		carry = (a & b) << 1
		a, b = sum, carry
		if carry == 0 {
			break
		}
	}
	return a
}

func SolveOffer65() {
	a, b := 1, 2
	fmt.Println(add1(a, b))
}
