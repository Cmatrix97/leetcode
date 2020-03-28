package interview

import "fmt"

// 将二进制各位求和，如果能整除3则为0，否则为1
func singleNumber1(nums []int) int {
	var bitSum [32]int
	for _, num := range nums {
		for i := 0; i < 32 && num != 0; i++ {
			bitSum[i] += num & 1
			num >>= 1
		}
	}
	var sum int
	for i := 0; i < len(bitSum); i++ {
		if bitSum[i]%3 != 0 {
			sum += 1 << i
		}
	}
	return sum
}

func SolveOffer56_2() {
	nums := []int{9, 1, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber1(nums))
}
