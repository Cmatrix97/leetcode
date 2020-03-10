package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin1_154(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

// @lc code=end

func Solve154() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Println(findMin1_154(nums))
}
