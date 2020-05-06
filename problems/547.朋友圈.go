/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 朋友圈
 */
package problems

import "fmt"

/*
dfs
*/
func findCircleNum1(M [][]int) int {
	visited := make([]bool, len(M))
	count := 0
	var dfs func(i int)
	dfs = func(i int) {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 1 && !visited[j] {
				visited[j] = true
				dfs(j)
			}
		}
	}
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}
	return count
}

// 并查集
// type UnionFind struct {
// 	parents []int
// 	size    []int
// 	count   int
// }

// func NewUF(n int) *UnionFind {
// 	var u UnionFind
// 	u.parents = make([]int, n)
// 	u.size = make([]int, n)
// 	for i := 0; i < n; i++ {
// 		u.parents[i] = i
// 		u.size[i] = 1
// 	}
// 	u.count = n
// 	return &u
// }

// func (u *UnionFind) union(node1, node2 int) {
// 	root1, root2 := u.find(node1), u.find(node2)
// 	if root1 == root2 {
// 		return
// 	}
// 	if u.size[root1] < u.size[root2] {
// 		u.parents[root1] = root2
// 		u.size[root2] += u.size[root1]
// 	} else {
// 		u.parents[root2] = root1
// 		u.size[root1] += u.size[root2]
// 	}
// 	u.count--
// }

// func (u *UnionFind) find(node int) int {
// 	for node != u.parents[node] {
// 		u.parents[node] = u.parents[u.parents[node]]
// 		node = u.parents[node]
// 	}
// 	return node
// }

// func (u *UnionFind) isConnected(node1, node2 int) bool {
// 	return u.find(node1) == u.find(node2)
// }

// func (u *UnionFind) getCount() int {
// 	return u.count
// }

func findCircleNum2(M [][]int) int {
	uf := NewUF(len(M))
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 1 {
				uf.union(i, j)
			}
		}
	}
	return uf.getCount()
}

// @lc code=start

// @lc code=end
func Solve547() {
	M := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(findCircleNum1(M))
}
