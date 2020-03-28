package interview

import (
	"container/list"
	"fmt"
	"math"
)

// 蛮力法
func maxSlidingWindow1(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var res []int
	for end := k; end <= len(nums); end++ {
		max := math.MinInt64
		for start := end - k; start < end; start++ {
			if nums[start] > max {
				max = nums[start]
			}
		}
		res = append(res, max)
	}
	return res
}

// 双端队列
func maxSlidingWindow2(nums []int, k int) []int {
	deque := list.New()
	var res []int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			deque.PushBack(0)
		}
		frontIdx := deque.Front().Value.(int)
		if i-frontIdx == k {
			deque.Remove(deque.Front())
		}
		var next *list.Element
		for node := deque.Front(); node != nil; node = next {
			next = node.Next()
			if nums[node.Value.(int)] <= nums[i] {
				deque.Remove(node)
			}
		}
		deque.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[deque.Front().Value.(int)])
		}
	}
	return res
}

func SolveOffer59_1() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow2(nums, k))
}
