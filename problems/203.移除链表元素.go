/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */
package problems

import "fmt"

/*迭代
增加空头结点
*/
func removeElements1(head *ListNode, val int) *ListNode {
	null := &ListNode{
		Next: head,
	}
	pre := null
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return null.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements2(head *ListNode, val int) *ListNode {
	null := &ListNode{}
	tail := null
	for head != nil {
		if head.Val != val {
			tail.Next = head
			tail = head
		} else {
			tail.Next = nil
		}
		head = head.Next
	}
	return null.Next
}

/*
递归写法，很好理解
*/
func removeElements3(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements3(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}

// @lc code=start

// @lc code=end

func Solve203() {
	values := []int{1, 2, 6, 3, 4, 5, 6}
	// values := []int{}
	val := 6
	head := CreateLinkedList(values)
	removeElements1(head, val)
	arrow := ""
	for head != nil {
		fmt.Print(arrow, head.Val)
		head = head.Next
		arrow = "->"
	}
}
