package problems

type Element interface{}

type Stack interface {
	Push(val int)
	Pop() int
	Peek() int
	Size() int
	IsEmpty() bool
}

var header *Entry
var size int

type Entry struct {
	Pre  *Entry
	Next *Entry
	El   Element
}

func newEntry(pre, next *Entry, el Element) *Entry {
	return &Entry{pre, next, el}
}

func NewStack() *Entry {
	header = newEntry(nil, nil, nil)
	header.Pre = header
	header.Next = header
	return header
}

func (e *Entry) Push(el Element) {
	newEntry := newEntry(header.Pre, header, el)
	newEntry.Pre.Next = newEntry
	newEntry.Next.Pre = newEntry
	size++
}

func (e *Entry) Pop() Element {
	if e.IsEmpty() {
		return nil
	}
	preEntry := header.Pre
	header.Pre = preEntry.Pre
	size--
	return preEntry.El
}

func (e *Entry) Peek() Element {
	if e.IsEmpty() {
		return nil
	}
	return header.Pre.El
}

func (e *Entry) Size() int {
	return size
}

func (e *Entry) IsEmpty() bool {
	if size != 0 {
		return false
	}
	return true
}
