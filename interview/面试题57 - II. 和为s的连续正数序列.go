package interview

import "fmt"

// 滑动窗口
func findContinuousSequence1(target int) [][]int {
	var res [][]int
	for start, end := 1, 2; start < end; {
		sum := (start + end) * (end - start + 1) / 2
		if sum > target {
			start++
		} else if sum < target {
			end++
		} else {
			var tmp []int
			for i := start; i <= end; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			end++
		}
	}
	return res
}

func SolveOffer57_2() {
	target := 15
	fmt.Println(findContinuousSequence1(target))
}
