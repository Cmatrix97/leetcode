package problem

import "math"

//Solution155 最小栈
func Solution155()  {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin());  // --> 返回 -3.
	minStack.Pop()
	println(minStack.Top())     // --> 返回 0.
	println(minStack.GetMin())   //--> 返回 -2.
}

/**
每次入栈先入最小值，再正常入栈；出栈执行两次
*/
type MinStack struct {
	stack []int
	top int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		top: 0,
	}
}


func (this *MinStack) Push(x int)  {
	min := math.MaxInt64
	if this.top == 0 {
		min = x
	} else {
		min = this.stack[this.top-2]
	}
	this.stack = append(this.stack,int(math.Min(float64(x),float64(min))))
	this.top++
	this.stack = append(this.stack, x)
	this.top++
}


func (this *MinStack) Pop()  {
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