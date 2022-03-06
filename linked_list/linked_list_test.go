package linked_list

import (
	"reflect"
	"testing"
)

func Test_StartsEmpty(t *testing.T) {
	ll := NewLinkedList()

	len := len(ll.ToSlice())
	if len != 0 {
		t.Errorf("expected new linked list to be empty, got len '%v'", len)
	}
}

func Test_Append(t *testing.T) {
	ll := NewLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(ll.ToSlice(), expected) {
		t.Errorf("expected added items to equal %v", expected)
	}
}

func Test_Prepend(t *testing.T) {
	ll := NewLinkedList()
	ll.Prepend(1)
	ll.Prepend(2)
	ll.Prepend(3)

	expected := []int{3, 2, 1}
	actual := ll.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected added items to equal %v, got %v", expected, actual)
	}
}

func Test_Delete(t *testing.T) {
	ll := NewLinkedList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ll.Delete(2)

	expected := []int{1, 3}
	actual := ll.ToSlice()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected added items to equal %v, got %v", expected, actual)
	}
}
