package problem

import "sort"

//Solution88 合并两个有序数组
func  Solution88()  {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}
	// nums1 := []int{0}
	// nums2 := []int{1}
	m, n := 3,3
	merge2(nums1, m, nums2, n)
	for _, v := range nums1 {
		print(v)
	}
}

/**先合并后排序
*/
func merge1(nums1 []int, m int, nums2 []int, n int)  {
	copy(nums1[m:], nums2[:])
	sort.Ints(nums1)
}

/**双指针
*/
func merge2(nums1 []int, m int, nums2 []int, n int)  {
	nums1Copy := make([]int, m)
	copy(nums1Copy[:], nums1[:m])
	p, p1, p2 := 0, 0, 0
	for p1 < m && p2 < n {
		if nums1Copy[p1] < nums2[p2] {
			nums1[p] = nums1Copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
		p++
	}
	if p1 < m {
		copy(nums1[p:],nums1Copy[p1:])
	}
	if p2 < n {
		copy(nums1[p:],nums2[p2:])
	}
}