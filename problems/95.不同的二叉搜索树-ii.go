/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees1(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	var generate func(start, end int) []*TreeNode
	generate = func(start, end int) []*TreeNode {
		var all_trees []*TreeNode
		if start > end {
			all_trees = append(all_trees, nil)
			return all_trees
		}
		for i := start; i <= end; i++ {
			left := generate(start, i-1)
			right := generate(i+1, end)
			for _, l := range left {
				for _, r := range right {
					curr := &TreeNode{
						Val:   i,
						Left:  l,
						Right: r,
					}
					all_trees = append(all_trees, curr)
				}
			}
		}
		return all_trees
	}
	return generate(1, n)
}

// @lc code=start

// @lc code=end
func Solve95() {
	n := 5
	res := generateTrees1(n)
	var levelOrder func(root *TreeNode, n int)
	levelOrder = func(root *TreeNode, n int) {
		var queue []*TreeNode
		queue = append(queue, root)
		fmt.Print(root.Val, " ")
		n--
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				fmt.Print(cur.Left.Val, " ")
				n--
				if n == 0 {
					break
				}
				queue = append(queue, cur.Left)
			} else {
				fmt.Print("null ")
			}
			if cur.Right != nil {
				fmt.Print(cur.Right.Val, " ")
				n--
				if n == 0 {
					break
				}
				queue = append(queue, cur.Right)
			} else {
				fmt.Print("null ")
			}
		}
	}
	for _, v := range res {
		levelOrder(v, n)
		fmt.Println()
	}
}
