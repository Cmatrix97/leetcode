/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */
package problems

import "fmt"

/*
将问题转化为K个鸡蛋扔T次最多可以测多少层楼
(1)K==1  N(K,T) = T+1
(2)T==1  N(K,T) = 2
(3)N(K,T) = N(K-1,T-1) + N(K,T-1)
*/
func superEggDrop1(K int, N int) int {
	var fun func(K, T int) int
	fun = func(K, T int) int {
		if K == 1 || T == 1 {
			return T + 1
		}
		return fun(K-1, T-1) + fun(K, T-1)
	}
	T := 1
	for fun(K, T) <= N {
		T++
	}
	return T
}

// @lc code=start

// @lc code=end
func Solve887() {
	K, N := 2, 100
	fmt.Println(superEggDrop1(K, N))
}
