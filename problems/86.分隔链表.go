/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */
package problems

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition1_86(head *ListNode, x int) *ListNode {
	beforeHead, afterHead := &ListNode{}, &ListNode{}
	before, after := beforeHead, afterHead
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next, before.Next = nil, afterHead.Next
	return beforeHead.Next
}

// @lc code=start

// @lc code=end

func Solve86() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}
	x := 3
	res := partition1_86(head, x)
	var arrow string
	for res != nil {
		fmt.Print(arrow, res.Val)
		arrow = "->"
		res = res.Next
	}
}
