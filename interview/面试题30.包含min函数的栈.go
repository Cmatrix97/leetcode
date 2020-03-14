package interview

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor_30() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	minVal := math.MaxInt64
	if len(this.min) != 0 {
		minVal = this.min[len(this.min)-1]
	}
	if x < minVal {
		minVal = x
	}
	this.min = append(this.min, minVal)
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func SolveOffer30() {
	obj := Constructor_30()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Println(obj.Min())
	obj.Pop()
	fmt.Println(obj.Top())
	fmt.Println(obj.Min())
}
