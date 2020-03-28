package interview

import "fmt"

// panic递归
func isBalanced1(root *TreeNode) (balance bool) {
	balance = true
	defer func() {
		if p := recover(); p != nil {
			balance = p.(bool)
		}
	}()
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left)
		right := depth(root.Right)
		if left-right > 1 || right-left > 1 {
			panic(false)
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	depth(root)
	return
}

func SolveOffer55_2() {
	arr := []int{3, 9, 20, -1, -1, 15, 7}
	root := new(TreeNode)
	CreateTree(root, arr)
	fmt.Println(isBalanced1(root))
}
