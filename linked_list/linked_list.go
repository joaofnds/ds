package linked_list

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(n int) *Node {
	return &Node{n, nil}
}

func (ll *LinkedList) Append(data int) {
	n := NewNode(data)

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

func (ll *LinkedList) Prepend(data int) {
	n := NewNode(data)
	n.next = ll.head
	ll.head = n
}

func (ll *LinkedList) Delete(data int) {
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

func (ll *LinkedList) ToSlice() []int {
	out := []int{}

	for curr := ll.head; curr != nil; curr = curr.next {
		out = append(out, curr.data)
	}

	return out
}
