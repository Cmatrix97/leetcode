package interview

// 虚拟头结点
func deleteNode1(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	pre, p := dummy, head
	for p.Val != val {
		pre, p = p, p.Next
	}
	pre.Next = p.Next
	return dummy.Next
}

// 信息交换法
// 若形参为 val *ListNode 可O(1)时间删除
// 与Next交换值，再删除Next即可
// 分三种情况
// (1)待删除结点不是尾结点
// (2)待删除=头结点=尾结点（只有一个结点）
// (3)链表中有多个结点且待删除结点为尾结点

func SolveOffer18() {
	arr := []int{4, 5, 1, 9}
	val := 4
	head := CreateLinkedList(arr)
	res := deleteNode1(head, val)
	PrintLinkedList(res)
}
