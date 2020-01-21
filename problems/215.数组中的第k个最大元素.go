/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package problems

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
直接调用标准库
*/
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/*
堆排序
*/
func findKthLargest2(nums []int, k int) int {
	N := len(nums)
	sink := func(a []int, k, N int) {
		for 2*k+1 <= N-1 {
			j := 2*k + 1
			if j+1 <= N-1 && a[j+1] > a[j] {
				j++
			}
			if a[k] >= a[j] {
				break
			}
			a[k], a[j] = a[j], a[k]
			k = j
		}
	}
	for i := N/2 - 1; i >= 0; i-- {
		sink(nums, i, N)
	}
	for i := 0; i < k; i++ {
		nums[0], nums[N-1] = nums[N-1], nums[0]
		N--
		sink(nums, 0, N)
	}
	return nums[N]
}

/*
实现heap.Interface接口
*/
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	var x interface{}
	x, *h = (*h)[len(*h)-1], (*h)[:len(*h)-1]
	return x
}

func findKthLargest3(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	for i := 0; i < k-1; i++ {
		heap.Pop(&h)
	}
	return heap.Pop(&h).(int)
}

// @lc code=start

// @lc code=end

func Solve215() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest1(nums, k))
}
