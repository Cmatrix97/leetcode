/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */
package problems

import "fmt"

/*
哈希表
*/
func majorityElement1_229(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	var res []int
	for k, v := range m {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}

/*
N/k的众数最多只有k-1个
*/
func majorityElement2_229(nums []int) []int {
	var c1, c2, v1, v2 int
	for _, c := range nums {
		if c == c1 {
			v1++
		} else if c == c2 {
			v2++
		} else {
			if v1 == 0 {
				c1, v1 = c, 1
			} else if v2 == 0 {
				c2, v2 = c, 1
			} else {
				v1--
				v2--
			}
		}
	}
	v1, v2 = 0, 0
	for _, c := range nums {
		if c == c1 {
			v1++
		} else if c == c2 {
			v2++
		}
	}
	var res []int
	if v1 > len(nums)/3 {
		res = append(res, c1)
	}
	if v2 > len(nums)/3 {
		res = append(res, c2)
	}
	return res
}

/*
N/k通用解法
遍历数组，对于每个元素
1.如果是候选人，该候选人票数+1
2.如果不是候选人
(1)若有候选人票数为0，则将其替换为当前元素
(2)若所有候选人都有票，则所有候选人票数-1
最后检查所有候选人是否满足条件
*/
func majorityElement3_229(nums []int) []int {
	var res []int
	solve := func(nums []int, k int) {
		candidates, votes := make([]int, k-1), make([]int, k-1)
		for _, v := range nums {
			flag := false
			for i, c := range candidates {
				if c == v {
					votes[i]++
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			for i := range candidates {
				if votes[i] == 0 {
					candidates[i], votes[i] = v, 1
					flag = true
					break
				}
			}
			if flag {
				continue
			}
			for i := range votes {
				votes[i]--
			}
		}
		for i := range votes {
			votes[i] = 0
		}
		for _, v := range nums {
			for i, c := range candidates {
				if c == v {
					votes[i]++
					break
				}
			}
		}
		for i := range votes {
			if votes[i] > len(nums)/k {
				res = append(res, candidates[i])
			}
		}
	}
	solve(nums, 3)
	return res
}

// @lc code=start

// @lc code=end
func Solve229() {
	nums := []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement1_229(nums))
}
