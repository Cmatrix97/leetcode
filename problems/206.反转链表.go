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
	var newHead *ListNode
	for head != nil {
		head, head.Next, newHead = head.Next, newHead, head
	}
	return newHead
}

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
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// @lc code=start

// @lc code=end

func Solve206() {
	nums := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(nums)
	res := reverseList1(head)
	arrow := ""
	for res != nil {
		fmt.Print(arrow, res.Val)
		res = res.Next
		arrow = "->"
	}
}
