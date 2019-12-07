/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */
package problems

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/*暴力法
 */
func threeSum1(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	m := make(map[string]bool)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					str := strings.Replace(strings.Trim(fmt.Sprint(temp), "[]"), " ", ",", -1)
					l := len(m)
					m[str] = true
					if len(m) == l {
						continue
					}
					res = append(res, temp)
				}
			}
		}
	}
	return res
}

// @lc code=start
/*排序双指针
 */
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if sum := nums[i] + nums[left] + nums[right]; sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}

// @lc code=end

/*参考某大神的Java代码
 */
func threeSum3(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	minValue, maxValue := math.MaxInt64, math.MinInt64
	negSize, posSize, zeroSize := 0, 0, 0
	for _, v := range nums {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
		if v > 0 {
			posSize++
		} else if v < 0 {
			negSize++
		} else {
			zeroSize++
		}
	}
	if zeroSize >= 3 {
		res = append(res, []int{0, 0, 0})
	}
	if negSize == 0 || posSize == 0 {
		return res
	}
	if minValue*2+maxValue > 0 {
		maxValue = -minValue * 2
	} else if maxValue*2+minValue < 0 {
		minValue = -maxValue * 2
	}
	m := make([]int, maxValue-minValue+1)
	negs, poss := []int{}, []int{}
	for _, v := range nums {
		if v < minValue || v > maxValue {
			continue
		}
		if m[v-minValue] == 0 {
			if v > 0 {
				poss = append(poss, v)
			} else if v < 0 {
				negs = append(negs, v)
			}
		}
		m[v-minValue]++
	}
	sort.Ints(poss)
	sort.Ints(negs)
	basej := 0
	for i := len(negs) - 1; i >= 0; i-- {
		nv := negs[i]
		minp := (-nv) >> 1
		for basej < len(poss) && poss[basej] < minp {
			basej++
		}
		for j := basej; j < len(poss); j++ {
			pv := poss[j]
			cv := -nv - pv
			if cv >= nv && cv <= pv {
				if cv == nv {
					if m[nv-minValue] > 1 {
						res = append(res, []int{nv, nv, pv})
					}
				} else if cv == pv {
					if m[pv-minValue] > 1 {
						res = append(res, []int{nv, pv, pv})
					}
				} else {
					if m[cv-minValue] > 0 {
						res = append(res, []int{nv, cv, pv})
					}
				}
			} else if cv < nv {
				break
			}
		}
	}
	return res
}

func Solution15() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	//nums := []int{1,2,-2,-1}
	//nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	for _, arr := range res {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
