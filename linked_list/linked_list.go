package linked_list

type node[T comparable] struct {
	data T
	next *node[T]
}

type linkedList[T comparable] struct {
	head *node[T]
}

func NewLinkedList[T comparable]() *linkedList[T] {
	return new(linkedList[T])
}

func newNode[T comparable](n T) *node[T] {
	return &node[T]{n, nil}
}

func (ll *linkedList[T]) Append(data T) {
	n := newNode(data)

	if ll.head == nil {
		ll.head = n
		return
	}

	curr := ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = n
}

func (ll *linkedList[T]) Prepend(data T) {
	n := newNode(data)
	n.next = ll.head
	ll.head = n
}

func (ll *linkedList[T]) Delete(data T) {
	if ll.head == nil {
		return
	}

	if ll.head.data == data {
		ll.head = ll.head.next
	}

	for curr := ll.head; curr.next != nil; curr = curr.next {
		if curr.next.data == data {
			curr.next = curr.next.next
			return
		}
	}
}

func (ll *linkedList[T]) ToSlice() []T {
	out := []T{}

	for curr := ll.head; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}

	return out
}
