/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
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
1.翻转两个链表
2.以第2题方法计算
3.再翻转得到最终结果
*/
func addTwoNumbers1_445(l1 *ListNode, l2 *ListNode) *ListNode {
	reverse := func(l *ListNode) *ListNode {
		head := &ListNode{}
		p := l
		for l != nil {
			l = l.Next
			p.Next = head.Next
			head.Next = p
			p = l
		}
		return head.Next
	}
	l1 = reverse(l1)
	l2 = reverse(l2)
	head := &ListNode{}
	p, q, r := l1, l2, head
	var carry int
	for p != nil || q != nil {
		sum := carry
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
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
	return reverse(head.Next)
}

/*
栈
*/
func addTwoNumbers2_445(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	head := &ListNode{}
	var carry int
	for len(stack1) != 0 || len(stack2) != 0 {
		sum := carry
		if len(stack1) != 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) != 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		node := &ListNode{
			Val: sum % 10,
		}
		carry = sum / 10
		node.Next, head.Next = head.Next, node
	}
	if carry != 0 {
		node := &ListNode{
			Val: carry,
		}
		node.Next, head.Next = head.Next, node
	}
	return head.Next
}

// @lc code=start
/*
递归
先遍历得到len1和len2
(1)如果len1和len2都为1,当前值应为(l1.val+l2.val)%10，进位为(l1.val+l2.val)/10
(2)如果len1>len2，递归计算(l1.next,l2)，当前值应为(l1.val+carry)%10，进位为/
(3)如果len1==len2，递归计算(l1.next,l2.next)，当前值应为(l1.val+l2.val+carry)，进位为/
*/
func addTwoNumbers3_445(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var add func(l1, l2 *ListNode, len1, len2 int) *ListNode
	add = func(l1, l2 *ListNode, len1, len2 int) *ListNode {
		var sum int
		if len1 == 1 && len2 == 1 {
			sum = l1.Val + l2.Val
		} else if len1 > len2 {
			l1.Next = add(l1.Next, l2, len1-1, len2)
			sum = l1.Val + carry
		} else {
			l1.Next = add(l1.Next, l2.Next, len1-1, len2-1)
			sum = l1.Val + l2.Val + carry
		}
		l1.Val, carry = sum%10, sum/10
		return l1
	}
	res1, res2 := l1, l2
	var len1, len2 int
	for l1 != nil {
		len1++
		l1 = l1.Next
	}
	for l2 != nil {
		len2++
		l2 = l2.Next
	}
	var res *ListNode
	if len1 > len2 {
		res = add(res1, res2, len1, len2)
	} else {
		res = add(res2, res1, len2, len1)
	}
	if carry != 0 {
		res1 = &ListNode{
			Val:  carry,
			Next: res,
		}
		return res1
	}
	return res
}

// @lc code=end

func Solve445() {
	l1 := &ListNode{Val: 7, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	res := addTwoNumbers1_445(l1, l2)
	var arrow string
	for res != nil {
		fmt.Print(arrow, res.Val)
		arrow = "->"
		res = res.Next
	}
}
