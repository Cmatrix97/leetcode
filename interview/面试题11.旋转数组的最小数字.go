package interview

import "fmt"

func minArray1(numbers []int) int {
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)>>1
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}

func SolveOffer11() {
	numbers := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray1(numbers))
}
