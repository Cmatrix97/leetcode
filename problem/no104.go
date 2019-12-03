package problem

import "math"

//Solution104 二叉树的最大深度
func Solution104()  {
	root := &TreeNode{
		Val : 3,
		Left: &TreeNode{
			Val : 9,
		},
		Right: &TreeNode{
			Val : 20,
			Left : &TreeNode{
				Val : 15,
			},
			Right : &TreeNode{
				Val: 7,
			},
		},	
	}
	print(maxDepth(root))
}

//TreeNode 二叉树节点
type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

/**DFS
左右子树最大高度+1
*/
 func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)),float64(maxDepth(root.Right)))) + 1
}