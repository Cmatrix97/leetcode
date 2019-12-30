/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
package problems

import "fmt"

import "sort"

/*
通过map去重
*/
var (
	// route [][]int
	// track []int
	map40 map[string]bool
)

func combinationSum21_40(candidates []int, target int) [][]int {
	map40 = make(map[string]bool)
	sort.Ints(candidates)
	backtrack1_40(candidates, target, 0)
	res := route
	route = nil
	return res
}

func backtrack1_40(candidates []int, target, start int) {
	if target == 0 {
		key := fmt.Sprint(track)
		if _, ok := map40[key]; ok {
			return
		}
		temp := make([]int, len(track))
		copy(temp, track)
		route = append(route, temp)
		map40[key] = true
		return
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		track = append(track, candidates[i])
		backtrack1_40(candidates, target-candidates[i], i+1)
		track = track[:len(track)-1]
	}
}

/*
通过剪枝去重
*/
// var (
// 	route [][]int
// 	track []int
// )

func combinationSum22_40(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	backtrack2_40(candidates, target, 0)
	res := route
	route = nil
	return res
}

func backtrack2_40(candidates []int, target, start int) {
	if target == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		route = append(route, temp)
		return
	}
	for i := start; i < len(candidates) && target-candidates[i] >= 0; i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		track = append(track, candidates[i])
		backtrack2_40(candidates, target-candidates[i], i+1)
		track = track[:len(track)-1]
	}
}

// @lc code=start

// @lc code=end
func Solve40() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum21_40(candidates, target)
	fmt.Println(res)
}
