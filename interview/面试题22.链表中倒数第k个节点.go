package interview

// 快慢双指针，提高了代码鲁棒性
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for k != 0 && fast != nil {
		fast = fast.Next
		k--
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	if slow == dummy {
		return nil
	}
	return slow
}

// 递归
var pos int

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	res := getKthFromEnd2(head.Next, k)
	pos++
	if k == pos {
		return head
	}
	return res
}

func SolveOffer22() {
	arr := []int{1, 2, 3, 4, 5}
	head := CreateLinkedList(arr)
	k := 2
	res := getKthFromEnd2(head, k)
	PrintLinkedList(res)
}
