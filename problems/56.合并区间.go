/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */
package problems

import (
	"fmt"
	"sort"
)

/*
分别排序区间上下限left和right
然后遍历，若left[i+1]<=right[i]，说明连续
直到left[i+1]>right[i]时该区间结束
*/
func merge1_56(intervals [][]int) [][]int {
	var left, right []int
	var res [][]int
	n := len(intervals)
	if n == 0 {
		return res
	}
	for i := range intervals {
		left = append(left, intervals[i][0])
		right = append(right, intervals[i][1])
	}
	sort.Ints(left)
	sort.Ints(right)
	left = append(left, right[n-1]+1)
	for i, start := 0, 0; i < n; i++ {
		if left[i+1] <= right[i] {
			continue
		}
		res = append(res, []int{left[start], right[i]})
		start = i + 1
	}
	return res
}

// @lc code=start
/*
只按照下限排序
*/
type interval [][]int

func (a interval) Len() int           { return len(a) }
func (a interval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a interval) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge(intervals [][]int) [][]int {
	var res [][]int
	if len(intervals) == 0 {
		return res
	}
	sort.Sort(interval(intervals))
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			if res[len(res)-1][1] <= intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

// @lc code=end

func Solve56() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	res := merge(intervals)
	fmt.Println(res)
}
