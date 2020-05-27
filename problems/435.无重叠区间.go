/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */
package problems

import (
	"fmt"
	"sort"
)

// 贪心 对右边界排序 优先选择较小的
func eraseOverlapIntervals1(intervals [][]int) int {
	si := sortedIntervals(intervals)
	sort.Sort(si)
	var count int
	var max int
	for i, interval := range intervals {
		if i == 0 || interval[0] >= max {
			count++
			max = interval[1]
		}
	}
	return len(intervals) - count
}

type sortedIntervals [][]int

func (si sortedIntervals) Len() int {
	return len(si)
}

func (si sortedIntervals) Less(i, j int) bool {
	return si[i][1] < si[j][1]
}

func (si sortedIntervals) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}

// @lc code=start

// @lc code=end
func Solve435() {
	intervals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}
	fmt.Println(eraseOverlapIntervals1(intervals))
}
