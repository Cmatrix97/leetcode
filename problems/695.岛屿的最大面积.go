/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */
package problems

import "fmt"

// dfs
func maxAreaOfIsland1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		if row < 0 || row >= M || col < 0 || col >= N {
			return 0
		}
		if grid[row][col] == 0 {
			return 0
		}
		grid[row][col] = 0
		sum := 1
		sum += dfs(row-1, col)
		sum += dfs(row+1, col)
		sum += dfs(row, col-1)
		sum += dfs(row, col+1)
		return sum
	}
	var max int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if count := dfs(i, j); count > max {
				max = count
			}
		}
	}
	return max
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

func maxAreaOfIsland2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	conv := func(c, r int) int {
		return c*N + r
	}
	island := func(c, r int) bool {
		if c < 0 || c >= M || r < 0 || r >= N {
			return false
		}
		if grid[c][r] == 0 {
			return false
		}
		return true
	}
	uf := NewUF(M * N)
	var hasLand bool
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				continue
			}
			hasLand = true
			if island(i+1, j) {
				uf.union(conv(i+1, j), conv(i, j))
			}
			if island(i, j+1) {
				uf.union(conv(i, j+1), conv(i, j))
			}
		}
	}
	if !hasLand {
		return 0
	}
	var max int
	for _, v := range uf.size {
		if v > max {
			max = v
		}
	}
	return max
}

// @lc code=start

// @lc code=end

func Solve695() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{1, 1, 0, 1, 1},
	}
	fmt.Println(maxAreaOfIsland1(grid))
}
