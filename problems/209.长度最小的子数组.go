/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */
package problems

import (
	"fmt"
	"math"
)

/*
构造前缀和数组
时间复杂度O(n^2)
*/
func minSubArrayLen1(s int, nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	sums := make([]int, N)
	sums[0] = nums[0]
	for i := 1; i < N; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	ans := math.MaxInt64
	for start := 0; start < N; start++ {
		for end := start; end < N; end++ {
			if sums[end]-sums[start]+nums[start] >= s {
				ans = int(math.Min(float64(ans), float64(end-start+1)))
				break
			}
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

/*
在方法1中遍历end时用二分查找找到第一个满足条件的
时间复杂度O(nlogn)
*/
func minSubArrayLen2(s int, nums []int) int {
	N := len(nums)
	if N == 0 {
		return 0
	}
	sums := make([]int, N)
	sums[0] = nums[0]
	for i := 1; i < N; i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	var binarySearch func(start, end, target int) (int, bool)
	binarySearch = func(start, end, target int) (int, bool) {
		var mid int
		for start <= end {
			mid = start + (end-start)>>1
			if sums[mid] < target {
				start = mid + 1
			} else if sums[mid] > target {
				end = mid - 1
			} else {
				return mid, true
			}
		}
		if sums[mid] > target {
			return mid, true
		}
		return 0, false
	}
	ans := math.MaxInt64
	for start := 0; start < N; start++ {
		target := sums[start] - nums[start] + s
		if end, ok := binarySearch(start, N-1, target); ok {
			ans = int(math.Min(float64(ans), float64(end-start+1)))
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

/*
双指针滑动窗口
时间复杂度O(n)
*/
func minSubArrayLen3(s int, nums []int) int {
	N := len(nums)
	ans := math.MaxInt64
	var left, right, sum int
	for right = 0; right < N; right++ {
		sum += nums[right]
		for sum >= s {
			ans = int(math.Min(float64(ans), float64(right-left+1)))
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// @lc code=start

// @lc code=end

func Solve209() {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	fmt.Println(minSubArrayLen1(s, nums))
}
