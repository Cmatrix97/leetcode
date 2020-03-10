package problems

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

// 对于每个节点，将其本身的值与左右子树结果相加作为该点为root的最大路径和
// 返回该点值与左右子树最大值的和作为该点为子树根节点的最大路径（不能分叉，只能左右选一条路）
func maxPathSum1(root *TreeNode) int {
	ans := math.MinInt64
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		L := max(dfs(root.Left), 0)
		R := max(dfs(root.Right), 0)
		ans = max(ans, L+R+root.Val)
		return max(L, R) + root.Val
	}
	dfs(root)
	return ans
}

// @lc code=end

func Solve124() {
	root := &TreeNode{}
	nums := []int{-10, 9, 20, -1, -1, 15, 7}
	InitTree(root, nums)
	res := maxPathSum1(root)
	fmt.Println(res)
}
