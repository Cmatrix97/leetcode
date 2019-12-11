/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */
package problems

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok && i-v <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

// @lc code=start

// @lc code=end

func Solve219() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
