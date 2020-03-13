package interview

import "fmt"

// 双指针
// 一个先序遍历，根左右
// 一个对称先序遍历，根右左
// 如果两指针所指节点值不相等则false
func isSymmetric1(root *TreeNode) bool {
	var compare func(p1, p2 *TreeNode) bool
	compare = func(p1, p2 *TreeNode) bool {
		if p1 == nil && p2 == nil {
			return true
		}
		if p1 == nil || p2 == nil || p1.Val != p2.Val {
			return false
		}
		return compare(p1.Left, p2.Right) && compare(p1.Right, p2.Left)
	}
	return compare(root, root)
}

func SolveOffer28() {
	root := new(TreeNode)
	arr := []int{1, 2, 2, 2, -1, 2}
	CreateTree(root, arr)
	fmt.Println(isSymmetric1(root))
}
