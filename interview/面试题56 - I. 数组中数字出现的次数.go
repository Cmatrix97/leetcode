package interview

import "fmt"

// map
// 时间复杂度O(n)
// 空间复杂度O(n)
func singleNumbers1(nums []int) []int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		} else {
			m[num] = struct{}{}
		}
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}

// 按两个数异或结果再进行分组异或
// 空间复杂度O(1)
func singleNumbers2(nums []int) []int {
	var xor, bit int
	for _, num := range nums {
		xor ^= num
	}
	for xor&1 == 0 {
		bit++
		xor >>= 1
	}
	div := 1 << bit
	var res []int
	res = append(res, 0, 0)
	for _, num := range nums {
		if num|div == num {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}

// 空间复杂度O(1)
func SovleOffer56_1() {
	nums := []int{1, 2, 5, 2}
	// nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers2(nums))
}
