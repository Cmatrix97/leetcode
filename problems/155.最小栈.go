/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
package problems

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	top   int
}

/** initialize your data structure here. */
func Constructor_155() MinStack {
	return MinStack{
		stack: []int{},
		top:   0,
	}
}

func (this *MinStack) Push(x int) {
	min := math.MaxInt64
	if this.top == 0 {
		min = x
	} else {
		min = this.stack[this.top-2]
	}
	this.stack = append(this.stack, int(math.Min(float64(x), float64(min))))
	this.top++
	this.stack = append(this.stack, x)
	this.top++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-2]
	this.top -= 2
}

func (this *MinStack) Top() int {
	return this.stack[this.top-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.top-2]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// @lc code=start

// @lc code=end

func Solve155() {
	minStack := Constructor_155()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // --> 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top())    // --> 返回 0.
	fmt.Println(minStack.GetMin()) //--> 返回 -2.
}
