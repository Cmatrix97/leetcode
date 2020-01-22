/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */
package problems

import "fmt"

/*
1.记录
(1)i的左积nums[0] * ... * nums[i-1]
(2)i的右积nums[i+1] * ... * nums[len(nums)-1]
2.左积*右积
*/
func productExceptSelf1(nums []int) []int {
	N := len(nums)
	L, R := make([]int, N), make([]int, N)
	L[0], R[N-1] = 1, 1
	for i := 1; i < N; i++ {
		L[i] = L[i-1] * nums[i-1]
	}
	for i := N - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = L[i] * R[i]
	}
	return res
}

/*
用nums和res分别保存左积和右积
*/
func productExceptSelf2(nums []int) []int {
	N := len(nums)
	res := make([]int, N)
	res[N-1] = 1
	for i := N - 2; i >= 0; i-- {
		res[i] = res[i+1] * nums[i+1]
	}
	var prev int
	prev, nums[0] = nums[0], 1
	for i := 1; i < N; i++ {
		temp := nums[i]
		nums[i] = nums[i-1] * prev
		res[i] = nums[i] * res[i]
		prev = temp
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve238() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf1(nums))
}
