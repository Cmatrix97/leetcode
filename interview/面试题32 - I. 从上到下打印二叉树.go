package interview

import "fmt"

func levelOrder1_32_1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}

func SolveOffer32_1() {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := new(TreeNode)
	CreateTree(root, arr)
	res := levelOrder1_32_1(root)
	fmt.Println(res)
}
