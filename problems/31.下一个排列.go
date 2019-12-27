/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
package problems

import (
	"fmt"
	"sort"
)

/*
从后向前扫描至数组起始位置，遇到第一个nums[i-1] < nums[i]时结束遍历
从后向前扫描至i，遇到第一个大于nums[i-1]，记下标为j
交换nums[i-1]与nums[j]
将nums[i:]按升序重新排列
*/
func nextPermutation1(nums []int) {
	var i, j int
	for i = len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	for j = len(nums) - 1; j >= i; j-- {
		if nums[j] > nums[i-1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
}

// @lc code=start

// @lc code=end

func Solve31() {
	nums := []int{1, 2, 3, 4, 5}
	nextPermutation1(nums)
	fmt.Println(nums)
}
