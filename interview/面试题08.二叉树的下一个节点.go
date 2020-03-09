package interview

import "fmt"

// (1)有Right节点 -> Right子树的最左节点
// (2)没有Right，自己是Left -> Parent
// (3)没有Right，自己是Right -> 向上一直找Parent，找到第一个自身是Left，返回该节点的Parent
func getNext1(node *PTreeNode) *PTreeNode {
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	if node.Parent.Left == node {
		return node.Parent
	}
	for node.Parent.Right == node {
		if node.Parent.Parent == nil {
			return nil
		}
		node = node.Parent
	}
	return node.Parent
}

//  3
// | \
// 9 20
//  /  \
// 15   7
func SolveOffer08() {
	node3 := &PTreeNode{Val: 3}
	node9 := &PTreeNode{Val: 9}
	node20 := &PTreeNode{Val: 20}
	node15 := &PTreeNode{Val: 15}
	node7 := &PTreeNode{Val: 7}

	node3.Left = node9
	node3.Right = node20
	node9.Parent = node3
	node20.Parent = node3
	node20.Left = node15
	node20.Right = node7
	node15.Parent = node20
	node7.Parent = node20

	next := getNext1(node3)
	fmt.Println(next.Val)
}
