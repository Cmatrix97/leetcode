/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package problems

import "fmt"

/*
并查集
*/
type UnionFind struct {
	parents []int
	count   int
}

func (u *UnionFind) init(totalNodes int) {
	u.parents = make([]int, totalNodes)
	for i := 0; i < totalNodes; i++ {
		u.parents[i] = i
	}
	u.count = totalNodes
}

func (u *UnionFind) union(node1, node2 int) {
	root1, root2 := u.find(node1), u.find(node2)
	if root1 != root2 {
		u.parents[root2] = root1
		u.count--
	}
}

func (u *UnionFind) find(node int) int {
	son := node
	for node != u.parents[node] {
		node = u.parents[node]
	}
	for son != node {
		u.parents[son], son = node, u.parents[son]
	}
	return node
}

func (u *UnionFind) isConnected(node1, node2 int) bool {
	return u.find(node1) == u.find(node2)
}

func (u *UnionFind) getCount() int {
	return u.count
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	directions := [][]int{
		{1, 0},
		{0, 1},
	}
	node := func(i, j int) int {
		return i*N + j
	}
	var uf UnionFind
	dummyNode := M * N
	uf.init(dummyNode + 1)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' {
				for _, direction := range directions {
					nextX, nextY := i+direction[0], j+direction[1]
					if nextX < M && nextY < N && grid[nextX][nextY] == '1' {
						uf.union(node(i, j), node(nextX, nextY))
					}
				}
			} else {
				uf.union(dummyNode, node(i, j))
			}
		}
	}
	return uf.getCount() - 1
}

// @lc code=start

// @lc code=end

func Solve200() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
