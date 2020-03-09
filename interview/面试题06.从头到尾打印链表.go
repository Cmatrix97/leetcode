package interview

import "fmt"

// stack
func reversePrint1(head *ListNode) []int {
	var stack, res []int
	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}
	for len(stack) != 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, v)
	}
	return res
}

// stack+原地reverse
func reversePrint2(head *ListNode) []int {
	var stack []int
	for ; head != nil; head = head.Next {
		stack = append(stack, head.Val)
	}
	for i, j := 0, len(stack) - 1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

// 递归系统栈
func reversePrint3(head *ListNode) []int {
	var res []int
	var recurse func(node *ListNode)
	recurse = func(node *ListNode) {
		if node == nil {
			return
		}
		recurse(node.Next)
		res = append(res, node.Val)
	}
	recurse(head)
	return res
}

func SolveOffer06() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	fmt.Println(reversePrint3(head))
}
