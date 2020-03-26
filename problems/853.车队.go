/*
 * @lc app=leetcode.cn id=853 lang=golang
 *
 * [853] 车队
 */
package problems

import (
	"fmt"
	"sort"
)

// 按剩余距离排序
// 计算每辆车到达终点所需时间duration
// 如果后面的车duration比前面短，就将duration设为前面车的duration
func carFleet1(target int, position []int, speed []int) int {
	groups := len(position)
	group := make([]car, len(position))
	for i := range group {
		group[i] = car{pos: position[i], duration: float64(target-position[i]) / float64(speed[i])}
	}
	sort.Sort(sortBy(group))
	for i := 1; i < len(group); i++ {
		if group[i-1].duration >= group[i].duration {
			group[i].duration = group[i-1].duration
			groups--
		}
	}
	return groups
}

type car struct {
	pos      int
	duration float64
}

type sortBy []car

func (s sortBy) Len() int {
	return len(s)
}

func (s sortBy) Less(i, j int) bool {
	return s[i].pos > s[j].pos
}

func (s sortBy) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// @lc code=start

// @lc code=end
func Solve853() {
	target := 10
	position := []int{6, 8}
	speed := []int{3, 2}
	fmt.Println(carFleet1(target, position, speed))
}
