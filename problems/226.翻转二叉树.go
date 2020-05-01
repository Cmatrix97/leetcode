/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */
package problems

import "fmt"

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

// @lc code=start

// @lc code=end

func Solve226() {
	nums := []int{4, 2, 7, 1, 3, 6, 9}
	root := &TreeNode{}
	InitTree(root, nums)
	root = invertTree(root)
	//7由根节点右孩子变为左孩子
	fmt.Println(root.Left.Val)
}
