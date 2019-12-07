/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */
package problems

import "fmt"

/*迭代
 */
func reverseList1(head *ListNode) *ListNode {
	newHead := &ListNode{}
	p := head
	for head != nil {
		head = head.Next
		p.Next = newHead.Next
		newHead.Next = p
		p = head
	}
	return newHead.Next
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 递归
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// @lc code=end

func Solution206() {
	nums := []int{1, 2, 3, 4, 5}
	null := &ListNode{}
	p := null
	for _, v := range nums {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		p.Next = node
		p = p.Next
	}
	res := reverseList(null.Next)
	arrow := ""
	for res != nil {
		fmt.Print(arrow, res.Val)
		res = res.Next
		arrow = "->"
	}
}
