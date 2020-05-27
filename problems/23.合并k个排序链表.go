/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
package problems

import (
	"container/heap"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 暴力法
func mergeKLists1(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for {
		minIdx := -1
		min := math.MaxInt64
		for i, list := range lists {
			if list == nil {
				continue
			}
			if list.Val < min {
				min = list.Val
				minIdx = i
			}
		}
		if minIdx == -1 {
			break
		}
		p.Next = lists[minIdx]
		p = p.Next
		lists[minIdx] = lists[minIdx].Next
		p.Next = nil
	}
	return dummy.Next
}
func mergeKLists2(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	var mh MinHeap
	for _, list := range lists {
		if list == nil {
			continue
		}
		mh = append(mh, list)
	}
	heap.Init(&mh)
	for len(mh) > 0 {
		p.Next = mh[0]
		p = p.Next
		mh[0] = mh[0].Next
		if mh[0] == nil {
			heap.Remove(&mh, 0)
		} else {
			heap.Fix(&mh, 0)
		}
		p.Next = nil
	}
	return dummy.Next
}

type MinHeap []*ListNode

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *MinHeap) Pop() interface{} {
	el := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return el
}

// @lc code=start

// @lc code=end
func Solve23() {
	nums := [][]int{
		{1, 4, 5},
		{1, 3, 4},
		{2, 6},
	}
	lists := make([]*ListNode, len(nums))
	for i := range lists {
		lists[i] = CreateLinkedList(nums[i])
	}
	head := mergeKLists1(lists)
	PrintLinkedList(head)
}
