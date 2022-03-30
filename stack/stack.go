package stack

type Node[T any] struct {
	data T
	prev *Node[T]
}

type Stack[T any] struct {
	head *Node[T]
}

func NewStack[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		return *new(T), false
	}

	return s.head.data, true
}

func (s *Stack[T]) Push(data T) {
	node := &Node[T]{data, s.head}
	s.head = node
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}

	s.head = s.head.prev
}

func (s *Stack[T]) ToSlice() []T {
	out := []T{}

	for curr := s.head; curr != nil; curr = curr.prev {
		out = append(out, curr.data)
	}

	return out
}
