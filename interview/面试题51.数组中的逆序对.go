package interview

import "fmt"

// 暴力法，超时35/139
func reversePairs1(nums []int) int {
	var total int
	m := make(map[int]int)
	for _, num := range nums {
		for k, v := range m {
			if num < k {
				total += v
			}
		}
		m[num]++
	}
	return total
}

// 归并排序，双指针从后向前
// 如果左边>右边，总数加上右边剩余数量
func reversePairs2(nums []int) int {
	aux := make([]int, len(nums))
	var total int
	sortAndCount := func(nums []int, lo, mi, hi int) {
		i, j := mi, hi
		for k := lo; k <= hi; k++ {
			aux[k] = nums[k]
		}
		for k := hi; k >= lo; k-- {
			if i < lo {
				nums[k] = aux[j]
				j--
			} else if j <= mi {
				nums[k] = aux[i]
				i--
			} else if aux[i] > aux[j] {
				total += j - mi
				nums[k] = aux[i]
				i--
			} else {
				nums[k] = aux[j]
				j--
			}
		}
	}
	var merge func(nums []int, lo, hi int)
	merge = func(nums []int, lo, hi int) {
		if hi <= lo {
			return
		}
		mi := lo + (hi-lo)>>1
		merge(nums, lo, mi)
		merge(nums, mi+1, hi)
		sortAndCount(nums, lo, mi, hi)
	}
	merge(nums, 0, len(nums)-1)
	return total
}

func SolveOffer51() {
	nums := []int{7, 5, 6, 4}
	fmt.Println(reversePairs2(nums))
}
