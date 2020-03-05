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

func SolveOffer06() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2}}}
	fmt.Println(reversePrint1(head))
}
