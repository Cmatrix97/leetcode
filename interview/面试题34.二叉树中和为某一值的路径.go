package interview

import "fmt"

func pathSum1(root *TreeNode, sum int) [][]int {
	var res [][]int
	var tmp []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum -= root.Val
		tmp = append(tmp, root.Val)
		if root.Left == nil && root.Right == nil && sum == 0 {
			arr := make([]int, len(tmp))
			copy(arr, tmp)
			res = append(res, arr)
		}
		dfs(root.Left)
		dfs(root.Right)
		sum += root.Val
		tmp = tmp[:len(tmp)-1]
	}
	dfs(root)
	return res
}

func SolveOffer34() {
	arr := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1, -1, -1}
	sum := 22
	root := new(TreeNode)
	CreateTree(root, arr)
	fmt.Println(pathSum1(root, sum))
}
