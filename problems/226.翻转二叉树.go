/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */
package problems

import "fmt"

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Left), invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// @lc code=end

func Solution226() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	root = invertTree(root)
	//7由根节点右孩子变为左孩子
	fmt.Println(root.Left.Val)
}
