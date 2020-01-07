/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
package problems

import (
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST1(root *TreeNode) bool {
	var recurse func(node *TreeNode, lower, upper int) bool
	recurse = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		val := node.Val
		if val <= lower || val >= upper {
			return false
		}
		if !recurse(node.Left, lower, val) {
			return false
		}
		if !recurse(node.Right, val, upper) {
			return false
		}
		return true
	}
	return recurse(root, math.MaxInt64, math.MaxInt64)
}

/*
非递归
*/
func isValidBST2(root *TreeNode) bool {
	var stack []*TreeNode
	var lowers, uppers []int
	update := func(root *TreeNode, lower, upper int) {
		stack = append(stack, root)
		lowers = append(lowers, lower)
		uppers = append(uppers, upper)
	}
	lower, upper := math.MinInt64, math.MaxInt64
	update(root, lower, upper)
	for len(stack) != 0 {
		root, lower, upper = stack[len(stack)-1], lowers[len(lowers)-1], uppers[len(uppers)-1]
		stack, lowers, uppers = stack[:len(stack)-1], lowers[:len(lowers)-1], uppers[:len(uppers)-1]
		if root == nil {
			continue
		}
		val := root.Val
		if val <= lower || val >= upper {
			return false
		}
		update(root.Left, lower, val)
		update(root.Right, val, upper)
	}
	return true
}

/*
中序遍历升序
*/
func isValidBST3(root *TreeNode) bool {
	var stack []*TreeNode
	preorder := math.MinInt64
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) != 0 {
			root, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if root.Val <= preorder {
				return false
			}
			preorder = root.Val
			root = root.Right
		}
	}
	return true
}

// @lc code=start

// @lc code=end

func Solve98() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(isValidBST1(root))
}
