package interview

import (
	"fmt"
	"math"
	"strconv"
)

// 找规律
// 1. n%100 < 10 => (n/100)*10
// 2. 10 <= n%100 < 20 => (n/100)*10 + n%100 - 10 + 1
// 3. n%100 >= 20 => (n/100)*10 + 10
func countDigitOne1(n int) int {
	count, pow := 0, 1
	for n >= pow {
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

// 递归
func countDigitOne2(n int) int {
	strN := strconv.Itoa(n)
	return numberOf1(strN)
}

func numberOf1(strN string) int {
	if len(strN) == 0 {
		return 0
	}
	first := int(strN[0] - '0')
	if len(strN) == 1 && first == 0 {
		return 0
	}
	if len(strN) == 1 && first > 0 {
		return 1
	}
	var numFirstDigit int
	if first > 1 {
		numFirstDigit = int(math.Pow10(len(strN) - 1))
	} else if first == 1 {
		post, _ := strconv.Atoi(strN[1:])
		numFirstDigit = post + 1
	}
	numOtherDigits := first * (len(strN) - 1) * int(math.Pow10(len(strN)-2))
	numRecursive := numberOf1(strN[1:])
	return numFirstDigit + numOtherDigits + numRecursive
}

func SolveOffer43() {
	n := 19
	fmt.Println(countDigitOne1(n))
}
