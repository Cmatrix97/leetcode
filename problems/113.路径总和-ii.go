/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
递归
*/
func pathSum1_113(root *TreeNode, sum int) [][]int {
	var res [][]int
	var temp []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum -= root.Val
		temp = append(temp, root.Val)
		if sum == 0 && root.Left == nil && root.Right == nil {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
		}
		dfs(root.Left)
		dfs(root.Right)
		sum += root.Val
		temp = temp[:len(temp)-1]
	}
	dfs(root)
	return res
}

// @lc code=start

// @lc code=end

func Solve113() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}
	sum := 22
	res := pathSum1_113(root, sum)
	fmt.Println(res)
}
