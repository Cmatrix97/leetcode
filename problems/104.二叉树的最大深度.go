/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */
package problems

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*DFS
左右子树最大高度+1
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

// @lc code=start

// @lc code=end

func Solve104() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := &TreeNode{}
	InitTree(root, nums)
	fmt.Print(maxDepth(root))
}
