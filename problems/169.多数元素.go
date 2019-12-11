/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package problems

import (
	"fmt"
	"sort"
)

/*哈希表
 */
func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	res, max := 0, 0
	for k, v := range m {
		if v > max {
			res, max = k, v
		}
	}
	return res
}

/*排序后中位数
 */
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*Boyer-Moore 投票算法
 */
func majorityElement3(nums []int) int {
	var candidate int
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// @lc code=start

// @lc code=end

func Solve169() {
	arr := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement1(arr))
}
