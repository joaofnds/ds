package stack

type Node struct {
	data int
	prev *Node
}

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.head.data, true
}

func (s *Stack) Push(data int) {
	node := &Node{data, s.head}
	s.head = node
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}

	s.head = s.head.prev
}

func (s *Stack) ToSlice() []int {
	out := []int{}

	for curr := s.head; curr != nil; curr = curr.prev {
		out = append(out, curr.data)
	}

	return out
}
