/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
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

/*
 非递归
 双指针
*/
func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	myHead := &ListNode{Next: head}
	front, cur := myHead, head
	for i := 1; i < m; i++ {
		front, cur = cur, cur.Next
	}
	prev := front
	for i := m; i <= n; i++ {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	front.Next.Next, front.Next = cur, prev
	return myHead.Next
}

/*
递归（回溯）
*/
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	left := head
	var stop bool
	var recurseAndReverse func(right *ListNode, m int, n int)
	recurseAndReverse = func(right *ListNode, m, n int) {
		if n == 1 {
			return
		}
		right = right.Next
		if m > 1 {
			left = left.Next
		}
		recurseAndReverse(right, m-1, n-1)
		if left == right || right.Next == left {
			stop = true
		}
		if !stop {
			left.Val, right.Val = right.Val, left.Val
			left = left.Next
		}
	}
	recurseAndReverse(head, m, n)
	return head
}

// @lc code=start

// @lc code=end

func Solve92() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	m, n := 2, 4
	head = reverseBetween1(head, m, n)
	var arrow string
	for head != nil {
		fmt.Print(arrow, head.Val)
		arrow = "->"
		head = head.Next
	}
}
