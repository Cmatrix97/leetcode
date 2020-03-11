package interview

import "fmt"

func hammingWeight1(n int) int {
	var count int
	s := fmt.Sprintf("%b", n)
	for _, v := range s {
		if v == '1' {
			count++
		}
	}
	return count
}

// n & n-1去除二进制最低位1
func hammingWeight2(n int) int {
	var count int
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func SolveOffer15() {
	n := 0b11111111111111111111111111111101
	fmt.Println(hammingWeight2(n))
}
