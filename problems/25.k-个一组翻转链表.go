/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */
package problems

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 双递归
func reverseKGroup1(head *ListNode, k int) *ListNode {
	var reverse func(head *ListNode) *ListNode
	reverse = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		newHead := reverse(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}
	temp := head
	for i := 1; i < k && temp != nil; i++ {
		temp = temp.Next
	}
	if temp == nil {
		return head
	}
	suf := temp.Next
	temp.Next = nil
	suf = reverseKGroup1(suf, k)
	new := reverse(head)
	head.Next = suf
	return new
}

// @lc code=start

// @lc code=end

func Solve25() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	head := CreateLinkedList(arr)
	head = reverseKGroup1(head, k)
	PrintLinkedList(head)
}
