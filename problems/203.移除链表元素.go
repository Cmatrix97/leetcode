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

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
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

// @lc code=end

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

func Solution203() {
	values := []int{1, 2, 6, 3, 4, 5, 6}
	// values := []int{}
	val := 6
	head := &ListNode{}
	p := head
	for _, v := range values {
		temp := &ListNode{
			Val: v,
		}
		p.Next = temp
		p = p.Next
	}
	head = head.Next
	removeElements1(head, val)
	arrow := ""
	for head != nil {
		fmt.Print(arrow, head.Val)
		head = head.Next
		arrow = "->"
	}
}
