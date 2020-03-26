package interview

import "fmt"

// 二分查找，找最左边的target，然后遍历
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return 0
	}
	var total int
	for i := left; i < len(nums) && nums[i] == nums[left]; i++ {
		total++
	}
	return total
}

// 直接二分查找-0.5和+0.5
func search2(nums []int, target int) int {
	binSearch := func(target float32) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + (right-left)>>1
			if float32(nums[mid]) < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	return binSearch(float32(target)+0.5) - binSearch(float32(target)-0.5)
}

func SolveOffer53_1() {
	nums := []int{1, 4}
	target := 4
	fmt.Println(search2(nums, target))
}
