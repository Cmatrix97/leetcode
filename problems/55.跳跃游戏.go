/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */
package problems

import (
	"fmt"
	"math"
)

/*
回溯
时间复杂度O(2^n)超时
*/
func canJump1(nums []int) bool {
	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		if index == len(nums)-1 {
			return true
		}
		end := int(math.Min(float64(len(nums)-1), float64(index+nums[index])))
		for i := end; i > index; i-- {
			if backtrack(i) == true {
				return true
			}
		}
		return false
	}
	return backtrack(0)
}

/*
贪心
*/
func canJump2(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

// @lc code=start

// @lc code=end

func Solve55() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump1(nums))
}
