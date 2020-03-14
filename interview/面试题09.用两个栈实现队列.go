package interview

import "fmt"

// (1)一个只管入，一个只管出
// 入O(1)出O(n)
// in:直接入栈stack1
// out:如果stack2空，stack1全部压入stack2
//     再出栈stack2
// (2)一个作队列，另一个辅助
// 入O(n)出O(1)
// in:如果stack1不空，stack1全部压入stack2
//    再入栈stack1
//    最后再将stack2全部压入stack1
// out:直接出栈stack1
type CQueue struct {
	in  Stack
	out Stack
}

func Constructor_09() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.out.Empty() {
		return this.out.Pop().(int)
	}
	for !this.in.Empty() {
		this.out.Push(this.in.Pop())
	}
	if this.out.Empty() {
		return -1
	}
	return this.out.Pop().(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func SolveOffer09() {
	obj := Constructor_09()
	obj.AppendTail(3)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
