package linked_list

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func NewLinkedList() *linkedList {
	return new(linkedList)
}

func newNode(n int) *node {
	return &node{n, nil}
}

func (ll *linkedList) Append(data int) {
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

func (ll *linkedList) Prepend(data int) {
	n := newNode(data)
	n.next = ll.head
	ll.head = n
}

func (ll *linkedList) Delete(data int) {
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

func (ll *linkedList) ToSlice() []int {
	out := []int{}

	for curr := ll.head; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}

	return out
}
