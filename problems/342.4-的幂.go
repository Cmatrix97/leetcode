/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */
package problems

import "fmt"

/*
4的幂转换为二进制一定为1000...其中0的数量为偶数个，转换为字符串判断
*/
func isPowerOfFour1(num int) bool {
	if num&(num-1) != 0 || num == 0 {
		return false
	}
	res := fmt.Sprintf("%b", num)
	if len(res)&1 == 0 {
		return false
	}
	return true
}

/*
4的幂和1010101010101010101010101010101作&运算为它本身
*/
func isPowerOfFour2(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 == num
}

// @lc code=start

// @lc code=end

func Solve342() {
	num := 16
	fmt.Println(isPowerOfFour1(num))
}
