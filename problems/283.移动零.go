/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */
package problems

import "fmt"

/*
从后向前遇到0将其后面全部前移，同时向末尾补0
*/
func moveZeroes1(nums []int) {
	j := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				nums[j-1] = nums[j]
			}
			nums[j] = 0
			j--
		}
	}
}

/*
双指针
*/
func moveZeroes2(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

/*
按顺序交换0与非0数的位置
*/
func moveZeroes3(nums []int) {
	zeroPoint := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroPoint = i
			break
		}
	}
	if zeroPoint == -1 {
		return
	}
	for i := zeroPoint; i+1 < len(nums); i++ {
		if nums[i+1] != 0 {
			nums[i+1], nums[zeroPoint] = nums[zeroPoint], nums[i+1]
			zeroPoint++
		}
	}
}

// @lc code=start

// @lc code=end

func Solve283() {
	nums := []int{0, 1, 0, 0, 12}
	moveZeroes1(nums)
	for _, v := range nums {
		fmt.Print(v, ",")
	}
}
