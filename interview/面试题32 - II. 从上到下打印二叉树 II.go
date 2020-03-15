package interview

import "fmt"

// 开一个额外的level队列记录层数
func levelOrder1_32_2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var tmp []int
	queue := []*TreeNode{root}
	level := []int{0}
	var maxDepth int
	for len(queue) > 0 {
		cur, depth := queue[0], level[0]
		queue, level = queue[1:], level[1:]
		if depth > maxDepth {
			maxDepth = depth
			arr := make([]int, len(tmp))
			copy(arr, tmp)
			res = append(res, arr)
			tmp = []int{}
		}
		tmp = append(tmp, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			level = append(level, depth+1)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			level = append(level, depth+1)
		}
	}
	res = append(res, tmp)
	return res
}

// 直接逐层遍历
func levelOrder2_32_2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tmp []int
		curLength := len(queue)
		for i := 0; i < curLength; i++ {
			cur := queue[0]
			queue = queue[1:]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

func SolveOffer32_2() {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := new(TreeNode)
	CreateTree(root, arr)
	res := levelOrder2_32_2(root)
	fmt.Println(res)
}
