/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] 被围绕的区域
 */
package problems

import "fmt"

/*
1.从边界的O开始dfs，将所有与其联通的O标记为#
2.将剩余的O全部置为X，将#置为O
递归DFS
*/
func solve1(board [][]byte) {
	if len(board) == 0 {
		return
	}
	M, N := len(board), len(board[0])
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= M || j >= N || board[i][j] == 'X' || board[i][j] == '#' {
			return
		}
		board[i][j] = '#'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if isEdge := i == 0 || j == 0 || i == M-1 || j == N-1; isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

/*
非递归DFS
*/
type Pos struct {
	x, y int
}

func solve2(board [][]byte) {
	if len(board) == 0 {
		return
	}
	direction := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	M, N := len(board), len(board[0])
	outOfBound := func(i, j int) bool {
		if i < 0 || j < 0 || i >= M || j >= N {
			return true
		}
		return false
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		var stack []*Pos
		stack = append(stack, &Pos{x: i, y: j})
		board[i][j] = '#'
	tag:
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			for i := 0; i < 4; i++ {
				nextX, nextY := top.x+direction[i][0], top.y+direction[i][1]
				if !outOfBound(nextX, nextY) && board[nextX][nextY] == 'O' {
					stack = append(stack, &Pos{x: nextX, y: nextY})
					board[nextX][nextY] = '#'
					continue tag
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if isEdge := i == 0 || j == 0 || i == M-1 || j == N-1; isEdge && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

/*
BFS
*/
// type Pos struct {
// 	x, y int
// }
func solve3(board [][]byte) {
	if len(board) == 0 {
		return
	}
	direction := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	M, N := len(board), len(board[0])
	outOfBound := func(i, j int) bool {
		if i < 0 || j < 0 || i >= M || j >= N {
			return true
		}
		return false
	}
	var bfs func(i, j int)
	bfs = func(i, j int) {
		var queue []*Pos
		queue = append(queue, &Pos{x: i, y: j})
		board[i][j] = '#'
		for len(queue) != 0 {
			cur := queue[0]
			queue = queue[1:]
			for i := 0; i < 4; i++ {
				nextX, nextY := cur.x+direction[i][0], cur.y+direction[i][1]
				if !outOfBound(nextX, nextY) && board[nextX][nextY] == 'O' {
					queue = append(queue, &Pos{x: nextX, y: nextY})
					board[nextX][nextY] = '#'
				}
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if isEdge := i == 0 || j == 0 || i == M-1 || j == N-1; isEdge && board[i][j] == 'O' {
				bfs(i, j)
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

// @lc code=start
/*
并查集Union Find
*/
// type UnionFind struct {
// 	parents []int
// }

// func (u *UnionFind) init(totalNodes int) {
// 	u.parents = make([]int, totalNodes)
// 	for i := range u.parents {
// 		u.parents[i] = i
// 	}
// }

// func (u *UnionFind) union(node1, node2 int) {
// 	root1, root2 := u.find(node1), u.find(node2)
// 	if root1 != root2 {
// 		u.parents[root2] = root1
// 	}
// }

// func (u *UnionFind) find(node int) int {
// 	son := node
// 	for u.parents[node] != node {
// 		node = u.parents[node]
// 	}
// 	for son != node {
// 		u.parents[son], son = node, u.parents[son]
// 	}
// 	return node
// }

// func (u *UnionFind) isConnected(node1, node2 int) bool {
// 	return u.find(node1) == u.find(node2)
// }

func solve4(board [][]byte) {
	if len(board) == 0 {
		return
	}
	M, N := len(board), len(board[0])
	direction := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}
	outOfBound := func(i, j int) bool {
		if i < 0 || j < 0 || i >= M || j >= N {
			return true
		}
		return false
	}
	var uf UnionFind
	uf.init(M*N + 1)
	dummyNode := M * N
	node := func(i, j int) int {
		return i*N + j
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == M-1 || j == 0 || j == N-1 {
					uf.union(dummyNode, node(i, j))
				} else {
					for k := 0; k < 4; k++ {
						nextX, nextY := i+direction[k][0], j+direction[k][1]
						if !outOfBound(nextX, nextY) && board[nextX][nextY] == 'O' {
							uf.union(node(i, j), node(nextX, nextY))
						}
					}
				}
			}
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if uf.isConnected(node(i, j), dummyNode) {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

// @lc code=end

func Solve130() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve1(board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%q ", board[i][j])
		}
		fmt.Println()
	}
}
