/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
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
 记录层数
*/
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	var levels, temp []int
	var level int
	levels = append(levels, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		level, levels = levels[0], levels[1:]
		root, queue = queue[0], queue[1:]
		fmt.Println(root.Val, level)
		if level == len(res)+1 {
			t := make([]int, len(temp))
			copy(t, temp)
			res = append(res, t)
			temp = []int{}
		}
		temp = append(temp, root.Val)
		if root.Left != nil {
			levels = append(levels, level+1)
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			levels = append(levels, level+1)
			queue = append(queue, root.Right)
		}
	}
	res = append(res, temp)
	return res
}

/*
记录宽度
*/
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		width := len(queue)
		var temp []int
		for i := 0; i < width; i++ {
			root, queue = queue[0], queue[1:]
			temp = append(temp, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve102() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := &TreeNode{}
	InitTree(root, nums)
	res := levelOrder(root)
	fmt.Println(res)
}
