package interview

import "fmt"

func levelOrder1_32_3(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var reverse bool
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
		if reverse {
			for i := 0; i < len(tmp)>>1; i++ {
				tmp[i], tmp[len(tmp)-1-i] = tmp[len(tmp)-1-i], tmp[i]
			}
		}
		res = append(res, tmp)
		reverse = !reverse
	}
	return res
}

func SolveOffer32_3() {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := new(TreeNode)
	CreateTree(root, arr)
	res := levelOrder1_32_3(root)
	fmt.Println(res)
}
