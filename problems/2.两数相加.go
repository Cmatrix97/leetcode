/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p, q, r := l1, l2, head
	var carry int
	for p != nil || q != nil {
		var pValue, qValue int
		if p != nil {
			pValue = p.Val
			p = p.Next
		}
		if q != nil {
			qValue = q.Val
			q = q.Next
		}
		sum := pValue + qValue + carry
		r.Next = &ListNode{
			Val: sum % 10,
		}
		carry = sum / 10
		r = r.Next
	}
	if carry != 0 {
		r.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

// @lc code=start

// @lc code=end

func Solve2() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := addTwoNumbers(l1, l2)
	var arrow string
	for res != nil {
		fmt.Print(arrow, res.Val)
		arrow = "->"
		res = res.Next
	}
}
