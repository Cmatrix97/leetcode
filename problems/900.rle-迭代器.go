/*
 * @lc app=leetcode.cn id=900 lang=golang
 *
 * [900] RLE 迭代器
 */
package problems

import "fmt"

type RLEIterator struct {
	data  []int
	count []int
}

func Constructor_900(A []int) RLEIterator {
	ri := RLEIterator{}
	for i := 0; i < len(A)>>1; i++ {
		ri.count = append(ri.count, A[i<<1])
		ri.data = append(ri.data, A[i<<1+1])
	}
	return ri
}

func (this *RLEIterator) Next(n int) int {
	for i := range this.count {
		if n <= this.count[i] {
			this.count[i] -= n
			return this.data[i]
		} else {
			n -= this.count[i]
			this.count[i] = 0
		}
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
// @lc code=start

// @lc code=end
func Solve900() {
	A := []int{3, 8, 0, 9, 2, 5}
	ri := Constructor_900(A)
	fmt.Println(ri.Next(2))
	fmt.Println(ri.Next(1))
	fmt.Println(ri.Next(1))
	fmt.Println(ri.Next(2))
}
