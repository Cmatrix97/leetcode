/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package problems

import (
	"fmt"
	"sort"
)

/*
回溯+剪枝
加法
*/
func combinationSum1_39(candidates []int, target int) [][]int {
	var route [][]int
	sort.Ints(candidates)
	track := make([]int, target/candidates[0]+1)
	backtrack1_39(candidates, track, target, 0, 0, &route)
	return route
}

func backtrack1_39(candidates, track []int, target, index, sum int, route *[][]int) {
	if sum == target {
		temp := make([]int, index)
		copy(temp, track)
		*route = append(*route, temp)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if index > 0 && candidates[i] < track[index-1] {
			continue
		}
		track[index] = candidates[i]
		if sum+track[index] > target {
			break
		}
		backtrack1_39(candidates, track, target, index+1, sum+track[index], route)
	}
}

/*
减法
定义为包级变量
*/
var (
	route [][]int
	track []int
)

func combinationSum2_39(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	backtrack2_39(candidates, target, 0)
	res := route
	route = nil
	return res
}

func backtrack2_39(candidates []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		route = append(route, temp)
		return
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		track = append(track, candidates[i])
		backtrack2_39(candidates, target-candidates[i], i)
		track = track[:len(track)-1]
	}
}

// @lc code=start

// @lc code=end

func Solve39() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum1_39(candidates, target)
	fmt.Println(res)
}
