/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */
package problems

import (
	"fmt"
	"sort"
)

func permuteUnique1(nums []int) [][]int {
	var route [][]int
	var track []int
	used := make([]bool, len(nums))
	sort.Ints(nums)
	var backtrack func()
	backtrack = func() {
		if len(track) == len(nums) {
			temp := make([]int, len(track))
			copy(temp, track)
			route = append(route, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			track = append(track, nums[i])
			backtrack()
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrack()
	return route
}

/*
交换元素
46题基础上剪枝去重
*/
func permuteUnique2(nums []int) [][]int {
	var route [][]int
	var backtrack func(first int)
	n := len(nums)
	backtrack = func(first int) {
		if first == n-1 {
			temp := make([]int, len(nums))
			copy(temp, nums)
			route = append(route, temp)
		}
		m := make(map[int]bool)
		for i := first; i < n; i++ {
			if _, ok := m[nums[i]]; ok {
				continue
			}
			m[nums[i]] = true
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}
	backtrack(0)
	return route
}

// @lc code=start

// @lc code=end
func Solve47() {
	nums := []int{1, 1, 2}
	res := permuteUnique1(nums)
	fmt.Println(res)
}
