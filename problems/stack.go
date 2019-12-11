package problems

type Element struct {
	next  *Element
	Value interface{}
}

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
