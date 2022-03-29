package queue_test

import (
	"ds/queue"
	"reflect"
	"testing"
)

func Test_QueueStartsEmpty(t *testing.T) {
	q := queue.NewQueue[int]()

	actual := q.ToSlice()
	if !reflect.DeepEqual(actual, []int{}) {
		t.Errorf("expected queue to start empty, got '%v'", actual)
	}
}

func Test_EnqueueAddsToTheEnd(t *testing.T) {
	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	actual := q.ToSlice()
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("queue should add to the end, expected %v, got %v", expected, actual)
	}
}

func Test_PeekReturnsTheFirstAddedElement(t *testing.T) {
	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)

	actual, ok := q.Peek()
	expected := 1

	if !ok {
		t.Error("should be ok")
	}

	if actual != expected {
		t.Errorf("after adding [1,2], peek should return %d, got '%v'", expected, actual)
	}
}

func Test_DequeueRemovesTheFirstElement(t *testing.T) {
	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Dequeue()

	expected := []int{2, 3}
	got := q.ToSlice()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func Test_DequeueReturnsFalseWhenQueueIsEmpty(t *testing.T) {
	q := queue.NewQueue[int]()
	v, ok := q.Dequeue()

	if ok {
		t.Error("calling Dequeue on an empty queue should give ok false")
	}

	if v != 0 {
		t.Errorf("calling Dequeue on an empty queue should return 0, got %d", v)
	}
}

func Test_IsEmptyAfterRemovingAddedElements(t *testing.T) {
	q := queue.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Dequeue()
	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("should be empty after enqueing and dequeueing twice")
	}
}
