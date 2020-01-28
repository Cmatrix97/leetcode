/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */
package problems

import "fmt"

/*
数据范围确定
桶排序
*/
func sortArray1(nums []int) []int {
	bucket := make([]int, 100001)
	for _, v := range nums {
		bucket[v+50000]++
	}
	var i int
	for idx, count := range bucket {
		for j := 0; j < count; j++ {
			nums[i] = idx - 50000
			i++
		}
	}
	return nums
}

// @lc code=start

// @lc code=end
func Solve912() {
	nums := []int{5, 1, 1, 2, 0, 0}
	res := sortArray1(nums)
	fmt.Println(res)
}
