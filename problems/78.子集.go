/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package problems

import "fmt"

/*
回溯法
*/
func subsets1(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	choose := make([]bool, n)
	var backtrack func(start int)
	backtrack = func(start int) {
		var temp []int
		for i, v := range choose {
			if v == true {
				temp = append(temp, nums[i])
			}
		}
		res = append(res, temp)
		for i := start; i < n; i++ {
			choose[i] = !choose[i]
			backtrack(i + 1)
			choose[i] = !choose[i]
		}
	}
	backtrack(0)
	return res
}

/*
位运算
幂集的个数为2^n
*/
func subsets2(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	for i := 0; i < 1<<n; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				temp = append(temp, nums[j])
			}
		}
		res = append(res, temp)
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve78() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets1(nums))
}
