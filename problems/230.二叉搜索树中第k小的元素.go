/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest1(root *TreeNode, k int) int {
	var res int
	var midOrder func(root *TreeNode)
	midOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		midOrder(root.Left)
		if k--; k == 0 {
			res = root.Val
			return
		}
		midOrder(root.Right)
	}
	midOrder(root)
	return res
}

/*
非递归
*/
func kthSmallest2(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if k--; k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// @lc code=start

// @lc code=end

func Solve230() {
	root := &TreeNode{}
	nums := []int{5, 3, 6, 2, 4, -1, -1, 1}
	k := 3
	InitTree(root, nums)
	fmt.Println(kthSmallest1(root, k))
}
