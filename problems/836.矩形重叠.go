/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */
package problems

import "fmt"

// 逆向思维 考虑不重叠的情况
func isRectangleOverlap1(rec1 []int, rec2 []int) bool {
	return !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0] || rec1[1] >= rec2[3] || rec2[1] >= rec1[3])
}

// 投影到x,y轴，可以通过判断max左边界小于min右边界确定线段相交
func isRectangleOverlap2(rec1 []int, rec2 []int) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	return max(rec1[0], rec2[0]) < min(rec1[2], rec2[2]) && max(rec1[1], rec2[1]) < min(rec1[3], rec2[3])
}

// @lc code=start

// @lc code=end

func Solve836() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap1(rec1, rec2))
}
