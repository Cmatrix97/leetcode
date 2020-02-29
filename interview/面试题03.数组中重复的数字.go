// 题目一：找出数组中重复的数字。
// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
// 请找出数组中任意一个重复的数字。
package interview

import (
	"fmt"
	"sort"
)

// 先排序再遍历
// T(n) = O(nlogn) 取决于排序算法
// S(n) = O(1)
func findRepeatNumber1(nums []int) int {
	sort.Ints(nums)
	for i := range nums {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 哈希表
// T(n) = O(n)
// S(n) = O(n)
func findRepeatNumber2(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v] = struct{}{}
	}
	return -1
}

// 遍历数组
// (1)nums[i] == i跳过
// (2)如果nums[i] == nums[nums[i]]说明找到了重复元素nums[i]
// (3)否则交换nums[i]和nums[nums[i]],让nums[i]到它应该在的位置
// (4)继续检查新的nums[i](原来的nums[nums[i]])
// T(n) = O(n)
// S(n) = O(1)
func findRepeatNumber3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

func SolveOffer03() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber3(nums))
}

// 题目二：不修改数组找出重复的数字
// 在一个长度为n+1的数组里所有数字都在1~n的范围内，所以数组中至少有一个数字是重复的。
// 请找出数组中任意一个重复的数字，但不能修改输入的数组。
// 例如，如果输入长度为8的数组{2,3,5,4,3,2,6,7}，那么对应的输出是重复数字2或者3。

// 二分查找
// T(n) = O(nlogn)
// S(n) = O(1)
func findRepeatNumberNoModify(nums []int) int {
	start, end := 1, len(nums)-1
	for start <= end {
		middle := start + (end-start)>>1
		count := countRange(nums, start, middle)
		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		if count > (middle - start + 1) {
			end = middle
		} else {
			start = middle + 1
		}
	}
	return -1
}

func countRange(nums []int, start, end int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}

func SolveOffer03_2() {
	nums := []int{2, 3, 5, 4, 3, 2, 6, 7}
	fmt.Println(findRepeatNumberNoModify(nums))
}
