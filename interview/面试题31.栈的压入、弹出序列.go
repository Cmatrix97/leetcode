package interview

import "fmt"

func validateStackSequences1(pushed []int, popped []int) bool {
	var stack []int
	for len(pushed) != 0 {
		stack = append(stack, pushed[0])
		pushed = pushed[1:]
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			popped, stack = popped[1:], stack[:len(stack)-1]
		}
	}
	return len(popped) == 0
}

func SolveOffer31() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences1(pushed, popped))
}
