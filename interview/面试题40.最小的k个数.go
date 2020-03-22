package interview

import (
	"container/heap"
	"fmt"
)

// 构造最小二叉堆，排序
func getLeastNumbers1(arr []int, k int) []int {
	N := len(arr) - 1
	var res []int
	for i := (N - 1) >> 1; i >= 0; i-- {
		sink(arr, i, N)
	}
	for i := 0; i < k; i++ {
		res = append(res, arr[0])
		arr[0], arr[N] = arr[N], arr[0]
		N--
		sink(arr, 0, N)
	}
	return res
}

func sink(arr []int, k, N int) {
	for 2*k+1 <= N {
		j := 2*k + 1
		if j < N && arr[j+1] < arr[j] {
			j++
		}
		if arr[j] > arr[k] {
			break
		}
		arr[j], arr[k] = arr[k], arr[j]
		k = j
	}
}

// 实现heap接口，最大堆
func getLeastNumbers2(arr []int, k int) []int {
	if k <= 0 {
		return nil
	}
	mh := new(MaxHeap)
	mh.Heap = append(mh.Heap, arr[:k]...)
	heap.Init(mh)
	for i := k; i < len(arr); i++ {
		if arr[i] < mh.Heap[0] {
			heap.Pop(mh)
			heap.Push(mh, arr[i])
		}
	}
	return mh.Heap
}

// 快排
func getLeastNumbers3(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return nil
	} else if k == len(arr) {
		return arr
	}
	quickSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSort(arr []int, lo, hi, k int) {
	m := partition(arr, lo, hi)
	if m == k {
		return
	} else if m > k {
		quickSort(arr, lo, m-1, k)
	} else {
		quickSort(arr, m+1, hi, k)
	}
}

func partition(arr []int, lo, hi int) int {
	temp := arr[lo]
	for lo < hi {
		for lo < hi && arr[hi] >= temp {
			hi--
		}
		arr[lo] = arr[hi]
		for lo < hi && arr[lo] <= temp {
			lo++
		}
		arr[hi] = arr[lo]
	}
	arr[lo] = temp
	return lo
}

// 计数排序
func getLeastNumbers4(arr []int, k int) []int {
	var nums [10001]int
	for _, v := range arr {
		nums[v]++
	}
	var res []int
	for i := 0; i < len(nums) && k != 0; {
		if nums[i] == 0 {
			i++
			continue
		}
		res = append(res, i)
		nums[i]--
		k--
	}
	return res
}

func SolveOffer40() {
	arr := []int{0, 1, 1, 2, 4, 4, 1, 3, 3, 2}
	k := 6
	res := getLeastNumbers1(arr, k)
	fmt.Println(res)
}
