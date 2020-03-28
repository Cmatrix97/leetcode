package interview

import (
	"container/list"
	"fmt"
)

type MaxQueue struct {
	queue *list.List
	deque *list.List
}

func Constructor59_2() MaxQueue {
	mq := MaxQueue{}
	mq.queue = list.New()
	mq.deque = list.New()
	return mq
}

func (this *MaxQueue) Max_value() int {
	if this.deque.Len() == 0 {
		return -1
	}
	return this.deque.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	var prev *list.Element
	for node := this.deque.Back(); node != nil; node = prev {
		prev = node.Prev()
		if node.Value.(int) < value {
			this.deque.Remove(node)
		}
	}
	this.deque.PushBack(value)
	this.queue.PushBack(value)
}

func (this *MaxQueue) Pop_front() int {
	if this.queue.Len() == 0 {
		return -1
	}
	x := this.queue.Remove(this.queue.Front()).(int)
	if x == this.deque.Front().Value.(int) {
		this.deque.Remove(this.deque.Front())
	}
	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func SolveOffer59_2() {
	mq := Constructor59_2()
	mq.Push_back(1)
	mq.Push_back(2)
	fmt.Println(mq.Max_value())
	fmt.Println(mq.Pop_front())
	fmt.Println(mq.Max_value())
}
