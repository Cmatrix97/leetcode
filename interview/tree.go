package interview

import "fmt"

// TreeNode:二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InitTree:根据int slice生成一颗二叉树
func InitTree(root *TreeNode, nums []int) {
	var queue []*TreeNode
	root.Val, nums = nums[0], nums[1:]
	queue = append(queue, root)
	for len(queue) != 0 {
		var cur *TreeNode
		cur, queue = queue[0], queue[1:]
		if nums[0] != -1 {
			cur.Left = &TreeNode{
				Val: nums[0],
			}
			queue = append(queue, cur.Left)
		}
		if nums = nums[1:]; len(nums) == 0 {
			break
		}
		if nums[0] != -1 {
			cur.Right = &TreeNode{
				Val: nums[0],
			}
			queue = append(queue, cur.Right)
		}
		if nums = nums[1:]; len(nums) == 0 {
			break
		}
	}
}

// 层次遍历
func LevelOrderReverse(root *TreeNode) {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Print(cur.Val, " ")
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}


// PTreeNode:含有父指针的二叉树节点
type PTreeNode struct {
	Val int
	Left  *PTreeNode
	Right *PTreeNode
	Parent *PTreeNode
}