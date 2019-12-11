/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */
package problems

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	stack1 *list.List
	stack2 *list.List
}

/** Initialize your data structure here. */
func Constructor_232() MyQueue {
	queue := MyQueue{}
	queue.stack1, queue.stack2 = &list.List{}, &list.List{}
	queue.stack1.Init()
	queue.stack2.Init()
	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	for this.stack2.Len() > 0 {
		el := this.stack2.Back()
		this.stack2.Remove(el)
		this.stack1.PushBack(el.Value)
	}
	this.stack1.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stack2.Len() > 0 {
		return this.stack2.Remove(this.stack2.Back()).(int)
	}
	for this.stack1.Len() > 0 {
		this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
	}
	return this.stack2.Remove(this.stack2.Back()).(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stack2.Len() > 0 {
		return this.stack2.Back().Value.(int)
	}
	for this.stack1.Len() > 0 {
		el := this.stack1.Back()
		this.stack1.Remove(el)
		this.stack2.PushBack(el.Value)

	}
	return this.stack2.Back().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.stack1.Len() != 0 || this.stack2.Len() != 0 {
		return false
	}
	return true
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

// @lc code=start

// @lc code=end

func Solve232() {
	queue := Constructor_232()

	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek()) //返回1
	queue.Push(3)
	fmt.Println(queue.Peek()) //返回1
	// fmt.Println(queue.Pop())   //返回1
	// fmt.Println(queue.Empty()) //返回false
}
