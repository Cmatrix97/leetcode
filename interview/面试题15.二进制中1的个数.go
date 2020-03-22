package interview

import "fmt"

// 打印直接数
func hammingWeight1(num uint32) int {
	var count int
	s := fmt.Sprintf("%b", num)
	for _, v := range s {
		if v == '1' {
			count++
		}
	}
	return count
}

// n & n-1去除二进制最低位1
func hammingWeight2(num uint32) int {
	var count int
	for num != 0 {
		num &= num - 1
		count++
	}
	return count
}
func SolveOffer15() {
	n := 0b00000000000000000000000000001011
	fmt.Println(hammingWeight2(uint32(n)))
}
