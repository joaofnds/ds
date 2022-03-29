package queue

type Node[T any] struct {
	data T
	next *Node[T]
}

type Queue[T any] struct {
	start *Node[T]
	end   *Node[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.start == nil
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.start == nil {
		return *new(T), false
	}

	return q.start.data, true
}

func (q *Queue[T]) Enqueue(data T) {
	node := &Node[T]{data, nil}
	if q.end != nil {
		q.end.next = node
	}
	q.end = node

	if q.start == nil {
		q.start = node
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		return *new(T), false
	}

	data := q.start.data

	q.start = q.start.next
	if q.start == nil {
		q.end = nil
	}

	return data, true
}

func (q *Queue[T]) ToSlice() []T {
	out := []T{}

	for curr := q.start; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}

	return out
}
