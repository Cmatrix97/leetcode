/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum1(root *TreeNode, sum int) (res bool) {
	defer func() {
		if r := recover(); r != nil {
			res = r.(bool)
		}
	}()
	if root == nil {
		return false
	}
	var dfs func(node *TreeNode, total int)
	dfs = func(node *TreeNode, total int) {
		if node.Left == nil && node.Right == nil {
			if total == sum {
				panic(true)
			}
			return
		}
		if node.Left != nil {
			dfs(node.Left, total+node.Left.Val)
		}
		if node.Right != nil {
			dfs(node.Right, total+node.Right.Val)
		}
	}
	dfs(root, root.Val)
	return
}

// @lc code=start
// @lc code=end

func Solve112() {
	nums := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}
	root := &TreeNode{}
	InitTree(root, nums)
	sum := 22
	fmt.Println(hasPathSum1(root, sum))
}
