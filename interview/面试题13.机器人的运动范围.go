package interview

import (
	"fmt"
)

// 并查集，该题不推荐
// func movingCount1(m int, n int, k int) int {
// 	var un UnionFind
// 	un.init(m * n)
// 	calc := func(i, j int) int {
// 		return i*n + j
// 	}
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if !isValid(i, j, k) {
// 				continue
// 			}
// 			if i > 0 && isValid(i-1, j, k) {
// 				un.union(calc(i-1, j), calc(i, j))
// 			} else if j > 0 && isValid(i, j-1, k) {
// 				un.union(calc(i, j-1), calc(i, j))
// 			}
// 		}
// 	}
// 	return un.getCount()
// }

// func isValid(i, j, k int) bool {
// 	var sum int
// 	for i != 0 || j != 0 {
// 		m, n := 0, 0
// 		if i != 0 {
// 			m = i % 10
// 			i /= 10
// 		}
// 		if j != 0 {
// 			n = j % 10
// 			j /= 10
// 		}
// 		sum += m + n
// 	}
// 	if sum <= k {
// 		return true
// 	}
// 	return false
// }

// type UnionFind struct {
// 	parents []int
// 	size    []int
// }

// func NewUF(N int) {
// 	u.parents = make([]int, N)
// 	u.size = make([]int, N)
// 	for i := 0; i < N; i++ {
// 		u.parents[i] = i
// 		u.size[i] = 1
// 	}
// }

// func (u *UnionFind) union(node1, node2 int) {
// 	root1, root2 := u.find(node1), u.find(node2)
// 	u.parents[root2] = root1
// 	u.size[root1] += u.size[root2]
// }

// func (u *UnionFind) find(node int) int {
// 	for node != u.parents[node] {
// 		u.parents[node] = u.parents[u.parents[node]]
// 		node = u.parents[node]
// 	}
// 	return node
// }

// func (u *UnionFind) getCount() int {
// 	return u.size[0]
// }

// dfs
func movingCount2(m int, n int, k int) int {
	calc := func(num int) int {
		var sum int
		for num != 0 {
			sum += num % 10
			num /= 10
		}
		return sum
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var count int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if visited[i][j] {
			return
		}
		if calc(i)+calc(j) <= k {
			visited[i][j] = true
			count++
			if i-1 >= 0 {
				dfs(i-1, j)
			}
			if i+1 < m {
				dfs(i+1, j)
			}
			if j-1 >= 0 {
				dfs(i, j-1)
			}
			if j+1 < n {
				dfs(i, j+1)
			}
		}
	}
	dfs(0, 0)
	return count
}

func SolveOffer13() {
	m, n, k := 3, 2, 17
	fmt.Println(movingCount2(m, n, k))
}
