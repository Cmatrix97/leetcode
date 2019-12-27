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
	backtrack(nums, track, used, &route)
	return route
}

func backtrack(nums, track []int, used []bool, route *[][]int) {
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
		backtrack(nums, track, used, route)
		track = track[:len(track)-1]
		used[i] = false
	}
}

// @lc code=start

// @lc code=end

func Solve46() {
	nums := []int{1, 2, 3}
	res := permute2(nums)
	fmt.Println(res)
}
