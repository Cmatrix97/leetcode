package interview

import (
	"container/heap"
	"fmt"
)

// 最大堆最小堆
type MedianFinder struct {
	minH MinHeap
	maxH MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	heap.Init(&mf.maxH)
	heap.Init(&mf.minH)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.maxH.Heap) == len(this.minH.Heap) {
		heap.Push(&this.maxH, num)
		heap.Push(&this.minH, heap.Pop(&this.maxH))
	} else {
		heap.Push(&this.minH, num)
		heap.Push(&this.maxH, heap.Pop(&this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxH.Heap) == len(this.minH.Heap) {
		return float64(this.maxH.Heap[0]+this.minH.Heap[0]) / 2
	} else {
		return float64(this.minH.Top().(int))
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func SolveOffer41() {
	mf := Constructor()
	mf.AddNum(1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
}
