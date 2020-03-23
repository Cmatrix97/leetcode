/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */
package problems

import (
	"fmt"
)

// 两次遍历
// 第一次生成max_left和max_right数组分别用于表示每个左右最高的柱子
// 第二次累计总和，如果当前柱子高度小于左右最高柱子中较小者，则sum += 较小柱子-当前柱子
// 时间复杂度 O(2n) = O(n)
// 空间复杂度 S(2n) = S(n)
func trap1(height []int) int {
	max_left := make([]int, len(height))
	max_right := make([]int, len(height))
	for i, j := 1, len(height)-2; i < len(height); i, j = i+1, j-1 {
		if height[i-1] > max_left[i-1] {
			max_left[i] = height[i-1]
		} else {
			max_left[i] = max_left[i-1]
		}
		if height[j+1] > max_right[j+1] {
			max_right[j] = height[j+1]
		} else {
			max_right[j] = max_right[j+1]
		}
	}
	var sum int
	for i, v := range height {
		var minSide int
		if max_left[i] < max_right[i] {
			minSide = max_left[i]
		} else {
			minSide = max_right[i]
		}
		if v < minSide {
			sum += minSide - v
		}
	}
	return sum
}

// @lc code=start
// 双指针一次遍历
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func trap2(height []int) int {
	var max_left, max_right, sum int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] > max_left {
				max_left = height[left]
			} else {
				sum += max_left - height[left]
			}
			left++
		} else {
			if height[right] > max_right {
				max_right = height[right]
			} else {
				sum += max_right - height[right]
			}
			right--
		}
	}
	return sum
}

// @lc code=end
func Solve42() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap2(height))
}
