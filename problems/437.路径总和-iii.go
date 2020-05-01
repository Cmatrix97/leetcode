/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */
package problems

import "fmt"

/*
双递归
第一次前序遍历每个节点，第二次求每个节点为根节点的可满足解
*/
func pathSum1_437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := computeNode(root, sum)
	res += pathSum1_437(root.Left, sum)
	res += pathSum1_437(root.Right, sum)
	return res
}

func computeNode(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += computeNode(root.Left, sum-root.Val)
	res += computeNode(root.Right, sum-root.Val)
	return res
}

/*
递归+数组
*/
func pathSum2_437(root *TreeNode, sum int) int {
	path := make([]int, 1000)
	return recurse(root, sum, path, 0)
}

func recurse(root *TreeNode, sum int, path []int, p int) int {
	if root == nil {
		return 0
	}
	temp := 0
	n := 0
	path[p] = root.Val
	for i := p; i >= 0; i-- {
		temp += path[i]
		if temp == sum {
			n++
		}
	}
	nl := recurse(root.Left, sum, path, p+1)
	nr := recurse(root.Right, sum, path, p+1)
	return n + nl + nr
}

/*
 匿名函数递归
*/
func pathSum3_437(root *TreeNode, sum int) int {
	var recurse func(root *TreeNode, path []int)
	var ans int

	recurse = func(root *TreeNode, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		temp := 0
		for i := len(path) - 1; i >= 0; i-- {
			temp += path[i]
			if temp == sum {
				ans++
			}
		}
		recurse(root.Left, path)
		recurse(root.Right, path)
	}

	recurse(root, make([]int, 0, 1000))
	return ans
}

/*
 回溯法+哈希表，思路同560题
*/
func pathSum4_437(root *TreeNode, sum int) int {
	m := make(map[int]int)
	return helper(root, 0, sum, m)
}

func helper(root *TreeNode, acc, target int, m map[int]int) int {
	if root == nil {
		return 0
	}
	count := 0
	acc += root.Val
	if acc == target {
		count++
	}
	if v, ok := m[acc-target]; ok {
		count += v
	}
	m[acc]++
	countLeft := helper(root.Left, acc, target, m)
	countRight := helper(root.Right, acc, target, m)
	m[acc]--
	return count + countLeft + countRight
}

// @lc code=start

// @lc code=end

func Solve437() {
	nums := []int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}
	root := &TreeNode{}
	InitTree(root, nums)
	sum := 8
	fmt.Println(pathSum1_437(root, sum))
}
