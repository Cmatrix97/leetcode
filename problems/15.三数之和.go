/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package problems

import (
	"fmt"
	"sort"
	"strings"
)

/*暴力法
 */
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	m := make(map[string]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					str := strings.Replace(strings.Trim(fmt.Sprint(temp), "[]"), " ", ",", -1)
					l := len(m)
					m[str] = true
					if len(m) == l {
						continue
					}
					res = append(res, temp)
				}
			}
		}
	}
	return res
}

/*排序双指针
 */
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if sum := nums[i] + nums[left] + nums[right]; sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve15() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	//nums := []int{1,2,-2,-1}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum1(nums)
	for _, arr := range res {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
