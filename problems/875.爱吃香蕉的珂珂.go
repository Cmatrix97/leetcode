/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */
package problems

import "fmt"

/*
暴力穷举
超时105/113
*/
func minEatingSpeed1(piles []int, H int) int {
	var speed int
	for {
		speed++
		var time int
		for _, v := range piles {
			time += v / speed
			if v%speed != 0 {
				time++
			}
		}
		if time <= H {
			return speed
		}
	}
}

/*
假设最好情况下所有pile%speed==0
speed从ceil(total/H)开始
*/
func minEatingSpeed2(piles []int, H int) int {
	ceil := func(x, y int) int {
		res := x / y
		if x%y != 0 {
			res++
		}
		return res
	}
	var total int
	for _, v := range piles {
		total += v
	}
	speed := ceil(total, H)
	for {
		var time int
		for _, v := range piles {
			time += ceil(v, speed)
		}
		if time <= H {
			return speed
		}
		speed++
	}
}

// @lc code=start

// @lc code=end
func Solve875() {
	piles := []int{3, 6, 7, 11}
	H := 8
	fmt.Println(minEatingSpeed1(piles, H))
}
