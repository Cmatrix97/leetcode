/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */
package problems

import "fmt"

/*
dfs
*/
func numIslands1(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= M || col >= N {
			return
		}
		if grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	var count int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

/*
bfs
*/
type node struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	var bfs func(row, col int)
	bfs = func(row, col int) {
		var queue []*node
		queue = append(queue, &node{row, col})
		grid[row][col] = '0'
		for len(queue) != 0 {
			row, col = queue[0].x, queue[0].y
			queue = queue[1:]
			if row-1 >= 0 && grid[row-1][col] == '1' {
				queue = append(queue, &node{row - 1, col})
				grid[row-1][col] = '0'
			}
			if row+1 < M && grid[row+1][col] == '1' {
				queue = append(queue, &node{row + 1, col})
				grid[row+1][col] = '0'
			}
			if col-1 >= 0 && grid[row][col-1] == '1' {
				queue = append(queue, &node{row, col - 1})
				grid[row][col-1] = '0'
			}
			if col+1 < N && grid[row][col+1] == '1' {
				queue = append(queue, &node{row, col + 1})
				grid[row][col+1] = '0'
			}
		}
	}
	var count int
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' {
				count++
				bfs(i, j)
			}
		}
	}
	return count
}

/*
UnionFind
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

func numIslands3(grid [][]byte) int {
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
	fmt.Println(numIslands1(grid))
}
