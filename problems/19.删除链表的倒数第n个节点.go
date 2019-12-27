/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */
package problems

/*
 双指针法
*/
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	myhead := &ListNode{Next: head}
	prev, post := myhead, myhead
	for i := 0; post.Next != nil; i++ {
		post = post.Next
		if i >= n {
			prev = prev.Next
		}
	}
	prev.Next = prev.Next.Next
	return myhead.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// @lc code=start

// @lc code=end

func Solve19() {
	head := &ListNode{Val: 1}
	n := 1
	removeNthFromEnd1(head, n)
}
