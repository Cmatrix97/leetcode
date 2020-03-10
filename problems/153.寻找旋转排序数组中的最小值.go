package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
// 二分
func findMin1_153(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// @lc code=end

func Solve153() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin1_153(nums))
}
