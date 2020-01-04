/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
package problems

import (
	"fmt"
	"sort"
)

/*
回溯
*/
func subsetsWithDup2(nums []int) [][]int {
	var res [][]int
	var temp []int
	sort.Ints(nums)
	var backtrack func(start int)
	backtrack = func(start int) {
		t := make([]int, len(temp))
		copy(t, temp)
		res = append(res, t)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			backtrack(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(0)
	return res
}

/*
位运算+map
*/
func subsetsWithDup1(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	m := make(map[string]bool)
	sort.Ints(nums)
	for i := 0; i < 1<<n; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				temp = append(temp, nums[j])
			}
		}
		key := fmt.Sprintf("%q", temp)
		if _, ok := m[key]; ok {
			continue
		}
		m[key] = true
		res = append(res, temp)
	}
	return res
}

// @lc code=start
/*
直接排序+位操作
2 2 2 2 2
1 1 0 0 0
1 0 1 0 0
0 1 1 0 0
0 1 0 1 0
0 0 0 1 1
可以观察出如果当前数字与前一个相同，且前一个未选择（为0）时，跳过
*/
func subsetsWithDup3(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < 1<<n; i++ {
		var temp []int
		var illegal bool
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				if j > 0 && nums[j] == nums[j-1] && i>>(j-1)&1 == 0 {
					illegal = true
					break
				} else {
					temp = append(temp, nums[j])
				}
			}
		}
		if !illegal {
			res = append(res, temp)
		}
	}
	return res
}

// @lc code=end

func Solve90() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup1(nums)
	fmt.Println(res)
}
