package interview

// 递归
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTree1(root.Left)
	right := mirrorTree1(root.Right)
	root.Left, root.Right = right, left
	return root
}

func SolveOffer27() {
	arr := []int{4, 2, 7, 1, 3, 6, 9}
	root := new(TreeNode)
	CreateTree(root, arr)
	res := mirrorTree1(root)
	LevelOrderReverse(res)
}
