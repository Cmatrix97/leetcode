/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */
package problems

import "fmt"

/*
暴力法，超时，时间复杂度O(n^3)
*/
func subarraySum1(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			sum := 0
			for i := start; i < end; i++ {
				sum += nums[i]
			}
			if sum == k {
				count++
			}
		}
	}
	return count
}

/*
记录累加和，时间复杂度O(n^2)
*/
func subarraySum2(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	count := 0
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	for start := 0; start < len(nums); start++ {
		for end := start + 1; end <= len(nums); end++ {
			if sum[end]-sum[start] == k {
				count++
			}
		}
	}
	return count
}

/*
哈希表存储{累加和，出现的次数}，时间复杂度O(n)
*/
func subarraySum3(nums []int, k int) int {
	count, sum := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}

// @lc code=start

// @lc code=end

func Solve560() {
	nums := []int{1, 1, 1}
	fmt.Println(subarraySum1(nums, 2))
}
