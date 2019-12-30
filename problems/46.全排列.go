/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
package problems

import (
	"fmt"
	"sort"
)

/*
基于31题实现
先将数组从小到大排序，然后遍历nextPermutation
*/
func permute1(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for {
		next := make([]int, len(nums))
		copy(next, nums)
		res = append(res, next)
		ok := nextPermutation(nums)
		if !ok {
			break
		}
	}
	return res
}

func nextPermutation(nums []int) bool {
	var i, j int
	for i = len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		return false
	}
	for j = len(nums) - 1; j >= i; j-- {
		if nums[j] > nums[i-1] {
			nums[i-1], nums[j] = nums[j], nums[i-1]
			break
		}
	}
	sort.Ints(nums[i:])
	return true
}

/*
回溯法
*/
func permute2(nums []int) [][]int {
	used := make([]bool, len(nums))
	track := make([]int, 0, len(nums))
	var route [][]int
	backtrack_46(nums, track, used, &route)
	return route
}

func backtrack_46(nums, track []int, used []bool, route *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, track)
		*route = append(*route, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		track = append(track, nums[i])
		backtrack_46(nums, track, used, route)
		track = track[:len(track)-1]
		used[i] = false
	}
}

/*
匿名函数
*/
func permute3(nums []int) [][]int {
	var route [][]int
	var track []int
	used := make([]bool, len(nums))
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
*/
func permute4(nums []int) [][]int {
	var route [][]int
	var backtrack func(first int)
	n := len(nums)
	backtrack = func(first int) {
		if first == n-1 {
			temp := make([]int, len(nums))
			copy(temp, nums)
			route = append(route, temp)
		}
		for i := first; i < n; i++ {
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

func Solve46() {
	nums := []int{1, 2, 3}
	res := permute4(nums)
	fmt.Println(res)
}
