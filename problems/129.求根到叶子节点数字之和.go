/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
/*
二叉树DFS
*/
func sumNumbers1(root *TreeNode) int {
	var ans, temp int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		temp = temp*10 + root.Val
		if root.Left == nil && root.Right == nil {
			ans += temp
		}
		dfs(root.Left)
		dfs(root.Right)
		temp /= 10
	}
	dfs(root)
	return ans
}

// @lc code=start

// @lc code=end
func Solve129() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	fmt.Println(sumNumbers1(root))
}
