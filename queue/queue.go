package queue

type Node struct {
	data int
	next *Node
}

type Queue struct {
	start *Node
	end   *Node
}

func (q *Queue) IsEmpty() bool {
	return q.start == nil
}

func (q *Queue) Peek() int {
	return q.start.data
}

func (q *Queue) Enqueue(data int) {
	node := &Node{data, nil}
	if q.end != nil {
		q.end.next = node
	}
	q.end = node

	if q.start == nil {
		q.start = node
	}
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	data := q.start.data

	q.start = q.start.next
	if q.start == nil {
		q.end = nil
	}

	return data, true
}

func (q *Queue) ToSlice() []int {
	out := []int{}

	for curr := q.start; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}

	return out
}
