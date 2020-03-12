package interview

import "math"

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p := dummy
	for l1 != nil || l2 != nil {
		v1, v2 := math.MaxInt64, math.MaxInt64
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		if v1 < v2 {
			p.Next = l1
			l1, p = l1.Next, p.Next
		} else {
			p.Next = l2
			l2, p = l2.Next, p.Next
		}
	}
	return dummy.Next
}

func SolveOffer25() {
	arr1 := []int{1, 2, 4}
	arr2 := []int{1, 3, 4}
	l1 := CreateLinkedList(arr1)
	l2 := CreateLinkedList(arr2)
	res := mergeTwoLists1(l1, l2)
	PrintLinkedList(res)
}
