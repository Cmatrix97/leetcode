/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */
package problems

import "fmt"

func removeDuplicates1_80(nums []int) int {
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[j-1] && nums[i] == nums[j-2] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	return j
}

// @lc code=start

// @lc code=end
func Solve80() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	// nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates1_80(nums))
	fmt.Println(nums)
}
