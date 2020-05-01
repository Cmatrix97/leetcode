/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	var recurse func(root *TreeNode)
	recurse = func(root *TreeNode) {
		if root == nil {
			return
		}
		recurse(root.Left)
		res = append(res, root.Val)
		recurse(root.Right)
	}
	recurse(root)
	return res
}

/*
迭代
如果有左孩子，则一直入栈
否则出栈当前结点并将右孩子入栈
*/
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) != 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}

// @lc code=start

// @lc code=end
func Solve94() {
	nums := []int{1, -1, 2, 3}
	root := &TreeNode{}
	InitTree(root, nums)
	res := inorderTraversal1(root)
	fmt.Println(res)
}
