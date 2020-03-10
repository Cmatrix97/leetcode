package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree1(root *TreeNode) int {
	var ans int
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		LDepth := dfs(root.Left) + 1
		RDepth := dfs(root.Right) + 1
		ans = max(ans, LDepth+RDepth)
		return max(LDepth, RDepth)
	}
	dfs(root)
	return ans
}

// @lc code=end

func Solve543() {
	root := &TreeNode{}
	nums := []int{1, 2, 3, 4, 5}
	InitTree(root, nums)
	res := diameterOfBinaryTree1(root)
	fmt.Println(res)
}
