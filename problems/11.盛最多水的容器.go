/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
package problems

import (
	"fmt"
	"math"
)

/*
暴力法
时间复杂度O(n^2)超时
*/
func maxArea1(height []int) int {
	var max int
	n := len(height)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			val := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if val > max {
				max = val
			}
		}
	}
	return max
}

/*
双指针法
*/
func maxArea2(height []int) int {
	var max int
	start, end := 0, len(height)-1
	for start < end {
		area := int(math.Min(float64(height[start]), float64(height[end]))) * (end - start)
		max = int(math.Max(float64(max), float64(area)))
		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return max
}

// @lc code=start

// @lc code=end

func Solve11() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea1(height))
}
