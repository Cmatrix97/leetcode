package interview

import "fmt"

// 左右双指针
func exchange1(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for i < j && nums[i]&1 == 1 {
			i++
		}
		for i < j && nums[j]&1 == 0 {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

// 快慢双指针
func exchange2(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

// 封装
func exchange3(nums []int) []int {
	for i, j := 0, 0; i < len(nums); i++ {
		if filter(nums[i]) {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func filter(n int) bool {
	if n&1 == 1 {
		return true
	}
	return false
}

func SolveOffer21() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange3(nums))
}
