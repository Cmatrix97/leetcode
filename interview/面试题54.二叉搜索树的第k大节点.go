package interview

import "fmt"

// flag递归
func kthLargest1(root *TreeNode, k int) int {
	var res int
	var find bool
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		if !find {
			inOrder(root.Right)
		}
		if !find && k == 1 {
			find = true
			res = root.Val
			return
		}
		k--
		if !find {
			inOrder(root.Left)
		}
	}
	inOrder(root)
	return res
}

// panic递归
func kthLargest2(root *TreeNode, k int) (res int) {
	defer func() {
		res = recover().(int)
	}()
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Right)
		if k == 1 {
			panic(root.Val)
		}
		k--
		inOrder(root.Left)
	}
	inOrder(root)
	return
}

func SolveOffer54() {
	arr := []int{5, 3, 6, 2, 4, -1, -1, 1}
	k := 3
	root := new(TreeNode)
	CreateTree(root, arr)
	fmt.Println(kthLargest2(root, k))
}
