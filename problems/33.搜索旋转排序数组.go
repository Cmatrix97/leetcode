/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package problems

import "fmt"

/*
两次二分查找
第一次找出旋转点，通过比较nums[mid]与nums[left]或nums[right]的大小
第二次通过比较target与nums[0]判断target在左段还是右段
*/
func search1(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	var mid int
	if nums[left] > nums[right] {
		for left <= right {
			mid = left + (right-left)/2
			if nums[mid] < nums[left] {
				right = mid
			} else if nums[mid] > nums[left] {
				left = mid
			} else {
				break
			}
		}
	} else {
		mid = length - 1
	}
	if target < nums[0] {
		return binarySearch(nums, mid+1, length-1, target)
	} else {
		return binarySearch(nums, 0, mid, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/*
一次二分查找
如果nums[mid] >= nums[left]说明左边升序，否则右边升序
考虑target是否在某一单调递增的区间内
*/
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// @lc code=start

// @lc code=end

func Solve33() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	// nums := []int{6, 7, 1, 2, 3, 4, 5}
	target := 0
	res := search1(nums, target)
	fmt.Println(res)
}
