/*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */
package problems

import (
	"fmt"
	"math/rand"
	"time"
)

func rand7() int {
	return rand.New(rand.NewSource(time.Now().Unix())).Intn(7) + 1
}

// 7进制
// 生成1-49的随机数，对41-49拒绝采样
func rand10_1() int {
	for {
		i, j := rand7()-1, rand7()-1
		total := i*7 + j + 1
		if total <= 40 {
			return total%10 + 1
		}
	}
}

// 合理利用拒绝的采样
// 第一次剩9个 9*7 = 63
// 第二次剩3个 3*7 = 21
// 第三次剩1个 重新循环
func rand10_2() int {
	for {
		i, j := rand7()-1, rand7()-1
		total := i*7 + j + 1
		if total <= 40 {
			return total%10 + 1
		}
		i, j = total-40-1, rand7()-1
		total = i*7 + j + 1
		if total <= 60 {
			return total%10 + 1
		}
		i, j = total-60-1, rand7()-1
		total = i*7 + j + 1
		if total <= 20 {
			return total%10 + 1
		}
	}
}

// @lc code=start

// @lc code=end

func Solve470() {
	fmt.Println(rand10_1())
}
