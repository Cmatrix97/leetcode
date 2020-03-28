package interview

import "fmt"

// 递归
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth1(root.Left)
	right := maxDepth1(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func SolveOffer55_1() {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := new(TreeNode)
	CreateTree(root, arr)
	fmt.Println(maxDepth1(root))
}
