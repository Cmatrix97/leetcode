/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */
package problems

import "fmt"

/*
1.先遍历nums1并插入到哈希表
2.遍历nums2时查询哈希表中是否有，有则输出并删除
*/
func intersection(nums1 []int, nums2 []int) []int {
	res, exists := []int{}, struct{}{}
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = exists
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve349() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	res := intersection(nums1, nums2)
	for _, v := range res {
		fmt.Print(v, " ")
	}
}
