package main

import (
	"ds/linked_list"
	"ds/queue"
	"ds/stack"

	"fmt"
)

func main() {
	ll()
	s()
	q()
}

func ll() {
	ll := linked_list.LinkedList{}
	ll.Append(2)
	ll.Append(3)
	ll.Prepend(1)
	ll.Delete(2)
	fmt.Printf("%+v\n", ll.ToSlice())
}

func s() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	fmt.Printf("%+v\n", s.ToSlice())
}

func q() {
	s := queue.Queue{}
	s.Enqueue(1)
	s.Enqueue(2)
	s.Enqueue(3)
	s.Dequeue()
	fmt.Printf("%+v\n", s.ToSlice())
}
