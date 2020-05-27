/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] 分糖果 II
 */
package problems

import (
	"fmt"
	"math"
)

func distributeCandies1_1103(candies int, num_people int) []int {
	ans := make([]int, num_people)
	count := 1
	for i := 0; ; i++ {
		idx := i % num_people
		if candies-count <= 0 {
			ans[idx] += candies
			return ans
		}
		ans[idx] += count
		candies -= count
		count++
	}
}

func distributeCandies2_1103(candies int, num_people int) []int {
	p := int(math.Sqrt(2*float64(candies)+0.25) - 0.5)
	R := int(float64(candies) - float64((p+1)*p)*0.5)
	rows, cols := p/num_people, p%num_people
	kids := make([]int, num_people)
	for i := range kids {
		kids[i] = (i+1)*rows + int(float64(rows*(rows-1))*0.5)*num_people
		if i < cols {
			kids[i] += i + 1 + rows*num_people
		}
	}
	kids[cols] += R
	return kids
}

// @lc code=start

// @lc code=end
func Solve1103() {
	candies, num_people := 10, 3
	ans := distributeCandies1_1103(candies, num_people)
	fmt.Println(ans)
}
