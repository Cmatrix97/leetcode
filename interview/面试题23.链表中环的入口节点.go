package interview

import "fmt"

// 找出链表中环的入口节点
// 先初始化两个指针slow和fast
// slow每次走一步，fast每次走两步
// 当slow == fast时停止，此时二者必然在环中
// fast继续每次走一步并计数，slow在原地等待
// 当二者相遇时停止，统计出环中节点个数count
// 让slow和fast重新指向head，fast向后走count步
// slow和fast同时每次向后走一步，相遇时即为环的入口
func entryNode1(head *ListNode) *ListNode {
	slow, fast := head, head
	for {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	count := 0
	for {
		fast = fast.Next
		count++
		if slow == fast {
			break
		}
	}
	slow, fast = head, head
	for count != 0 {
		fast = fast.Next
		count--
	}
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}

func SolveOffer23() {
	arr := []int{1, 2, 3, 4, 5, 6}
	head := CreateLinkedList(arr)
	p, q := head, head
	for p.Next != nil {
		p = p.Next
	}
	for i := 0; i < 2; i++ {
		q = q.Next
	}
	p.Next = q
	// PrintLinkedList(head)
	fmt.Println(entryNode1(head).Val)
}
