package interview

// 堆
type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h Heap) Top() interface{} {
	return h[0]
}

// 最小堆
type MinHeap struct {
	Heap
}

func (mh MinHeap) Less(i, j int) bool {
	return mh.Heap[i] < mh.Heap[j]
}

// 最大堆
type MaxHeap struct {
	Heap
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh.Heap[i] > mh.Heap[j]
}
