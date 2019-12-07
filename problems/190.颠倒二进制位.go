/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */
package problems

import "fmt"

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = (res << 1) | (num & 1)
		num = num >> 1
	}
	return res
}

// @lc code=end

func Solution190() {
	var num uint32 = 0b00000010100101000001111010011100
	res := reverseBits(num)
	fmt.Print(res)
}
