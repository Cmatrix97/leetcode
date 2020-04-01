package interview

import (
	"fmt"
	"sort"
)

func isStraight1(nums []int) bool {
	sort.Ints(nums)
	var zero int
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			zero++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
		zero -= nums[i+1] - nums[i] - 1
	}
	return zero >= 0
}

func SolveOffer61() {
	nums := []int{1, 2, 0, 4, 5}
	fmt.Println(isStraight1(nums))
}
