package interview

import "fmt"

// 双指针法
// 设不重合段长度分别为x和y，重合段为z
// x + z + y
// y + z + x
// 最终必然相遇（或者没有公共节点都为nil）
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	nodeA, nodeB := headA, headB
	for nodeA != nodeB {
		if nodeA != nil {
			nodeA = nodeA.Next
		} else {
			nodeA = headB
		}
		if nodeB != nil {
			nodeB = nodeB.Next
		} else {
			nodeB = headA
		}
	}
	return nodeA
}

// 栈
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	var stackA, stackB []*ListNode
	for headA != nil || headB != nil {
		if headA != nil {
			stackA = append(stackA, headA)
			headA = headA.Next
		}
		if headB != nil {
			stackB = append(stackB, headB)
			headB = headB.Next
		}
	}
	var nodeCount int
	pA, pB := len(stackA), len(stackB)
	for pA > 0 && pB > 0 && stackA[pA-1] == stackB[pB-1] {
		pA--
		pB--
		nodeCount++
	}
	if nodeCount == 0 {
		return nil
	}
	return stackA[pA]
}

func SolveOffer52() {
	numsA := []int{4, 1, 8, 4, 5}
	numsB := []int{5, 0, 1}
	// numsB := []int{5, 0, 1, 8, 4, 5}
	headA := CreateLinkedList(numsA)
	headB := CreateLinkedList(numsB)
	pA, pB := headA, headB
	for pA.Val != 8 {
		pA = pA.Next
	}
	for pB.Val != 1 {
		pB = pB.Next
	}
	pB.Next = pA
	res := getIntersectionNode1(headA, headB)
	fmt.Println(res)
}
