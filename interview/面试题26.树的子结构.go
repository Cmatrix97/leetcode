package interview

import "fmt"

// 两次dfs
// 第一次搜索A中与B根节点相同值的节点
// 第二次对两个节点向下dfs
func isSubStructure1(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}
	rootA := searchRoot(A, B)
	if rootA == nil {
		return false
	}
	return judge(rootA, B)
}

func searchRoot(A *TreeNode, B *TreeNode) *TreeNode {
	if A == nil || A.Val == B.Val {
		return A
	}
	if lA := searchRoot(A.Left, B); lA != nil {
		return lA
	}
	if rA := searchRoot(A.Right, B); rA != nil {
		return rA
	}
	return nil
}

func judge(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return judge(A.Left, B.Left) && judge(A.Right, B.Right)
}

// 搜索的同时进行判断
func isSubStructure2(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	var res bool
	if A.Val == B.Val {
		res = judge(A, B)
	}
	if !res {
		res = isSubStructure2(A.Left, B)
	}
	if !res {
		res = isSubStructure2(A.Right, B)
	}
	return res
}

func SolveOffer26() {
	arrA := []int{3, 4, 5, 1, 2}
	arrB := []int{4, 1}
	rootA, rootB := new(TreeNode), new(TreeNode)
	CreateTree(rootA, arrA)
	CreateTree(rootB, arrB)
	fmt.Println(isSubStructure2(rootA, rootB))
}
