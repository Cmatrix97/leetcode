/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */
package problems

import "fmt"

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

/*
1.遍历树直到找到p和q
保存每个结点的父节点到map即parent[k] = v
2.将p的全部祖先放入祖先集合
3.迭代查询q的父节点直到发现与p的某一祖先相同
*/
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var stack []*TreeNode
	parent := make(map[*TreeNode]*TreeNode)
	stack = append(stack, root)
	for len(stack) != 0 {
		var node *TreeNode
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
		_, okP := parent[p]
		_, okQ := parent[q]
		if okP && okQ {
			break
		}
	}
	ancestors := make(map[*TreeNode]bool)
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}
	for k, v := range parent {
		fmt.Println(k, v)
	}
	for k, v := range ancestors {
		fmt.Println(k, v)
	}
	for {
		if _, ok := ancestors[q]; ok {
			break
		}
		q = parent[q]
	}
	return q
}

/*
递归
left表示左子树是否有p或q
right表示左子树是否有p或q
mid表示当前结点是否为p或q
三者有其二表示该结点为最近公共祖先
*/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var recurse func(root *TreeNode) bool
	recurse = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		var left, mid, right int
		if recurse(root.Left) {
			left = 1
		}
		if recurse(root.Right) {
			right = 1
		}
		if root == p || root == q {
			mid = 1
		}
		if left+mid+right == 2 {
			ans = root
		}
		return left+mid+right > 0
	}
	recurse(root)
	return ans
}

/*
递归
以root为根的树是否有p或q
(1)有p则返回p
(2)有q则返回q
(3)都有则返回root
(4)都没有返回nil
*/
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil && right == nil {
		return left
	} else if left == nil && right != nil {
		return right
	} else {
		return nil
	}
}

// @lc code=start

// @lc code=end

func Solve236() {
	root := &TreeNode{}
	nums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	InitTree(root, nums)
	var locate func(root *TreeNode, val int) *TreeNode
	locate = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return nil
		}
		if root.Val == val {
			return root
		}
		if left := locate(root.Left, val); left != nil {
			return left
		}
		if right := locate(root.Right, val); right != nil {
			return right
		}
		return nil
	}
	p, q := locate(root, 5), locate(root, 1)
	res := lowestCommonAncestor1(root, p, q)
	fmt.Println(res.Val)
}
