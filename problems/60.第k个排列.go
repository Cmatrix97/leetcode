/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */
package problems

import (
	"bytes"
	"fmt"
	"sort"
)

/*
利用nextPermutation不断生成下一个排列
*/
func getPermutation1(n int, k int) string {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	nextPermutation := func() {
		var i, j int
		for i = len(nums) - 1; i > 0; i-- {
			if nums[i-1] < nums[i] {
				break
			}
		}
		for j = len(nums) - 1; j >= i; j-- {
			if nums[j] > nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				break
			}
		}
		sort.Ints(nums[i:])
	}
	for i := 0; i < k-1; i++ {
		nextPermutation()
	}
	var buf bytes.Buffer
	for _, v := range nums {
		fmt.Fprint(&buf, v)
	}
	return buf.String()
}

/*
数学方法
1234
1243
1324
1342
1423
1432

2134
2143
2314
2341
2413
2431

3124
3142

3214
3241

3412
3421 18

4123
4132
4213
4231
4312
4321

每组(n-1)!个
(k-1)/(n-1)! + 1 = res
------------------------
如n = 4, k = 18, k-1 = 17:
(18-1)/(4-1)! + 1 = 3 为第一位
k-1 = (k-1) % (n-1)! = 17 % 6 = 5
n = n - 1 = 3
------------------------
5/(3-1)! + 1 = 3 为第二位，3被选过，所以为4
k-1 = 5 % 2 = 1
n = n - 1 = 2
------------------------
1/(2-1)! + 1 = 2 为第三位
k-1 = 1 % 1 = 0
n = n - 1 = 1
------------------------
0/(1-1)! + 1 = 1 为第四位
*/
func getPermutation2(n int, k int) string {
	fac := func(n int) int {
		res := 1
		for i := n; i > 0; i-- {
			res *= i
		}
		return res
	}
	var buf bytes.Buffer
	seen := make([]bool, n+1)
	k -= 1
	for i := n; i > 0; i-- {
		th := k/fac(i-1) + 1
		var j int
		for j = 1; j <= n; j++ {
			if seen[j] {
				continue
			}
			th--
			if th == 0 {
				seen[j] = true
				break
			}
		}
		fmt.Fprint(&buf, j)
		k %= fac(i - 1)
	}
	return buf.String()
}

// @lc code=start

// @lc code=end
func Solve60() {
	n, k := 4, 18
	fmt.Println(getPermutation1(n, k))
}
