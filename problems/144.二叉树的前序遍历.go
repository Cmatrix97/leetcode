/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
递归
*/
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	var preorder func(root *TreeNode)
	preorder = func (root *TreeNode)  {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)
	return res
}

/*
非递归
*/
func preorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		res = append(res, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve144() {
	nums := []int{1, -1, 2, 3}
	root := &TreeNode{}
	InitTree(root, nums)
	res := preorderTraversal1(root)
	fmt.Println(res)
}
