/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
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
*/
func swapPairs1(head *ListNode) *ListNode {
	myhead := &ListNode{Next: head}
	for prev := myhead; prev.Next != nil && prev.Next.Next != nil; {
		prev, prev.Next, prev.Next.Next, prev.Next.Next.Next = prev.Next, prev.Next.Next, prev.Next.Next.Next, prev.Next
	}
	return myhead.Next
}

/*
递归
*/
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

// @lc code=start

// @lc code=end

func Solve24() {
	arr := []int{1, 2, 3, 4}
	head := CreateLinkedList(arr)
	res := swapPairs1(head)
	var arrow string
	for res != nil {
		fmt.Print(arrow, res.Val)
		res = res.Next
		arrow = "->"
	}
}
