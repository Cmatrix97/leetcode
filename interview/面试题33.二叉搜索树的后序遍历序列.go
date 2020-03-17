package interview

import "fmt"

func verifyPostorder1(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}
	lastIdx := len(postorder) - 1
	cur := postorder[lastIdx]
	var split int
	for split = lastIdx - 1; split >= 0; split-- {
		if postorder[split] < cur {
			break
		}
	}
	for j := 0; j < split; j++ {
		if postorder[j] > cur {
			return false
		}
	}
	return verifyPostorder1(postorder[:split+1]) && verifyPostorder1(postorder[split+1:lastIdx])
}

func SolveOffer33() {
	arr := []int{4, 8, 6, 12, 16, 14, 10}
	fmt.Println(verifyPostorder1(arr))
}
