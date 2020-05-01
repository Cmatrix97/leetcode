/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
102题基础上交替逆序
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var flag bool
	reverse := func(temp []int) {
		for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
			temp[i], temp[j] = temp[j], temp[i]
		}
	}
	for len(queue) != 0 {
		flag = !flag
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
		if flag == false {
			reverse(temp)
		}
		res = append(res, temp)
	}
	return res
}

// @lc code=start

// @lc code=end

func Solve103() {
	nums := []int{3, 9, 20, -1, -1, 15, 7}
	root := &TreeNode{}
	InitTree(root, nums)
	res := zigzagLevelOrder(root)
	fmt.Println(res)
}
