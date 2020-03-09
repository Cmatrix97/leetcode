package problems

import "fmt"

// ListNode:单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateLinkedList:根据数组创建链表
func CreateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	secondNode := CreateLinkedList(arr[1:])
	head := &ListNode{
		Val:  arr[0],
		Next: secondNode,
	}
	return head
}

// PrintLinkedList:顺序输出链表
func PrintLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Print("nil")
}
