/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */
package problems

import (
	"fmt"
	"sort"
)

/*
min{糖果数量的一半,糖果种类数}
先排序然后遍历一次
*/
func distributeCandies1(candies []int) int {
	count := 1
	sort.Ints(candies)
	for i := 1; i < len(candies); i++ {
		if count == len(candies)/2 {
			return count
		}
		if candies[i] != candies[i-1] {
			count++
		}
	}
	return count
}

/*
通过set记录种类数量
*/
func distributeCandies2(candies []int) int {
	m := make(map[int]struct{})
	for i := 0; i < len(candies); i++ {
		if len(m) == len(candies)/2 {
			break
		}
		m[candies[i]] = struct{}{}
	}
	return len(m)
}

// @lc code=start

// @lc code=end

func Solve575() {
	candies := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies1(candies))
}
