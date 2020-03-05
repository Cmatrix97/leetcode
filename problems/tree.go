package problems

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InitTree 根据int slice生成一颗二叉树
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
