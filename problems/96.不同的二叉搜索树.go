/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */
package problems

import "fmt"

/*
动态规划
G(n):长度为n的序列的不同二叉搜索树个数（所求）
F(i,n):长度为n的序列中以i为根的不同二叉搜索树个数
则G(n) = F(1,n)+F(2,n)+...+F(n,n)
F(i,n) = G(i-1)*G(n-i) （左右子树情况的笛卡尔积）
*/
func numTrees1(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

/*
卡特兰数
C[n+1] = (4n+2)/n+2 * C[n]
*/
func numTrees2(n int) int {
	c := 1
	for i := 1; i < n; i++ {
		c = c * (4*i + 2) / (i + 2)
	}
	return c
}

// @lc code=start

// @lc code=end
func Solve96() {
	n := 5
	fmt.Println(numTrees1(n))
}
