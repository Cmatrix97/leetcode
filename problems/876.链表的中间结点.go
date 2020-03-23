/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */
package problems

import "fmt"

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢双指针
func middleNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		return slow.Next
	}
	return slow
}

// @lc code=end
func Solve876() {
	arr := []int{1, 2, 3, 4, 5}
	list := CreateLinkedList(arr)
	res := middleNode1(list)
	fmt.Println(res.Val)
}
