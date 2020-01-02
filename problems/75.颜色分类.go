/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */
package problems

import (
	"fmt"
)

/*
两次扫描
分别计数
*/
func sortColors1(nums []int) {
	var red, white, blue int
	for _, v := range nums {
		switch v {
		case 0:
			red++
		case 1:
			white++
		case 2:
			blue++
		}
	}
	for i := 0; i < red; i++ {
		nums[i] = 0
	}
	for i := red; i < red+white; i++ {
		nums[i] = 1
	}
	for i := red + white; i < len(nums); i++ {
		nums[i] = 2
	}
}

/*
三指针一次遍历
p0表示0的最右边界，p2表示2的最右边界，cur表示当前考虑的坐标
(1)nums[cur] = 0，交换下标为cur和p0的元素，并将指针都右移（因为在p0在cur左边，已经考虑过）
(2)nums[cur] = 2，交换下标为cur和p2的元素，将p2左移（p2在cur右侧，还没考虑过，所以cur不动）
(3)nums[cur] = 1，cur右移
*/
func sortColors(nums []int) {
	cur, p0, p2 := 0, 0, len(nums)-1
	for cur <= p2 {
		switch nums[cur] {
		case 0:
			nums[p0], nums[cur] = nums[cur], nums[p0]
			p0++
			cur++
		case 1:
			cur++
		case 2:
			nums[p2], nums[cur] = nums[cur], nums[p2]
			p2--
		}
	}
}

// @lc code=start

// @lc code=end

func Solve75() {
	nums := []int{2, 0, 1}
	sortColors(nums)
	fmt.Println(nums)
}
