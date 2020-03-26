package interview

import "fmt"

// 二分
func missingNumber1(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 求和
func missingNumber2(nums []int) int {
	total := len(nums) * (len(nums) + 1) >> 1
	for _, num := range nums {
		total -= num
	}
	return total
}

// 异或
func missingNumber3(nums []int) int {
	ans := len(nums)
	for i, num := range nums {
		ans ^= i ^ num
	}
	return ans
}

func SolveOffer53_2() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber3(nums))
}
