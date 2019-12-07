/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */
package problems

import "fmt"

/*
生成末尾0的唯一途径是2 * 5
而5的数量必然比2少
因此所求可转化为求5的数量
*/
func trailingZeroes1(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		for cur := i; (cur%5 == 0) && (cur/5 > 0); {
			count++
			cur /= 5
		}
	}
	return count
}

// @lc code=start
/*
n/5 + n/25 + n/125 + ...
*/
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n /= 5
	}
	return count
}

// @lc code=end

func Solution172() {
	n := 5
	fmt.Println(trailingZeroes(n))
}
