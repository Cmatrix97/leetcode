/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
dfs中序遍历
*/
func rightSideView1(root *TreeNode) []int {
	res := make([]int, 20)
	if root == nil {
		return res[:0]
	}
	level, maxLevel := -1, -1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		level++
		if level > maxLevel {
			maxLevel = level
		}
		dfs(root.Left)
		res[level] = root.Val
		dfs(root.Right)
		level--
	}
	dfs(root)
	return res[:maxLevel+1]
}

/*
dfs根右左
*/
func rightSideView2(root *TreeNode) []int {
	var res []int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level >= len(res) {
			res = append(res, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 0)
	return res
}

/*
bfs层次遍历记录每层最后一个
*/
func rightSideView3(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		width := len(queue)
		for i := 0; i < width; i++ {
			root, queue = queue[0], queue[1:]
			if i == width-1 {
				res = append(res, root.Val)
			}
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve199() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	res := rightSideView1(root)
	fmt.Println(res)
}
