/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */
package problems

import "fmt"

/*
循环右移并取最低位
*/
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1)
	}
	return count
}

// @lc code=end

func Solution191() {
	var num uint32 = 0b11110000111
	fmt.Println(hammingWeight(num))
}
