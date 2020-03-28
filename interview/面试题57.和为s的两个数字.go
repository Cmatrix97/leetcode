package interview

import "fmt"

// 双指针
func twoSum1_57(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		if sum := nums[lo] + nums[hi]; sum > target {
			hi--
		} else if sum < target {
			lo++
		} else {
			return []int{nums[lo], nums[hi]}
		}
	}
	return []int{}
}

func SolveOffer57() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum1_57(nums, target))
}
