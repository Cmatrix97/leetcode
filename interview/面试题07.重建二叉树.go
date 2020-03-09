package interview

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 分治思想
// preorder [3,9,20,15,7]
// inorder [9, 3, 15,20,7]
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	cur := &TreeNode{Val: preorder[0]}
	var part int
	for i, v := range inorder {
		if v == cur.Val {
			part = i
			break
		}
	}
	cur.Left = buildTree1(preorder[1:1+part], inorder[:part])
	cur.Right = buildTree1(preorder[1+part:], inorder[part+1:])
	return cur
}

func SolveOffer07() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree1(preorder, inorder)
	LevelOrderReverse(root)
}
