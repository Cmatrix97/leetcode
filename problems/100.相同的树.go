/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree1(p.Left, q.Left) && isSameTree1(p.Right, q.Right)
}

// @lc code=start
// @lc code=end

func Solve100() {
	numsP := []int{1, 2, 3}
	numsQ := []int{1, 2, 3}
	p, q := &TreeNode{}, &TreeNode{}
	InitTree(p, numsP)
	InitTree(q, numsQ)

	fmt.Println(isSameTree1(p, q))
}
