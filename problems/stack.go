package problems

type Element struct {
	next  *Element
	Value interface{}
}

/**
链栈
*/
type Stack struct {
	top Element
}

func (s *Stack) Push(v interface{}) {
	e := &Element{Value: v}
	e.next = s.top.next
	s.top.next = e
}

func (s *Stack) Pop() interface{} {
	e := s.top.next
	s.top.next = e.next
	return e.Value
}

func (s *Stack) Peek() interface{} {
	return s.top.next.Value
}

func (s *Stack) Empty() bool {
	if s.top.next != nil {
		return false
	}
	return true
}

/**
顺序栈
*/
type SStack []interface{}

func (s *SStack) Push(v interface{}) {
	*s = append(*s, v)
}

func (s *SStack) Pop() interface{} {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}

func (s SStack) Peek() interface{} {
	return s[len(s)-1]
}

func (s SStack) Empty() bool {
	return len(s) == 0
}
