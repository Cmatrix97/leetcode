package interview

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}

func SolveOffer24() {
	arr := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(arr)
	res := reverseList1(head)
	PrintLinkedList(res)
}
