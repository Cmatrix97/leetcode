package interview

import "fmt"

// 找规律
// 1. n%100 < 10 => (n/100)*10
// 2. 10 <= n%100 < 20 => (n/100)*10 + n%100 - 10 + 1
// 3. n%100 >= 20 => (n/100)*10 + 10
func countDigitOne1(n int) int {
	var count int
	pow := 1
	for n/pow != 0 {
		count += n / (pow * 10) * pow
		if incr := n%(pow*10) - pow + 1; incr > pow {
			count += pow
		} else if incr > 0 {
			count += incr
		}
		pow *= 10
	}
	return count
}

func SolveOffer43() {
	n := 19
	fmt.Println(countDigitOne1(n))
}
