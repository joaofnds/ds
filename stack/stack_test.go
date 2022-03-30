package stack_test

import (
	"ds/stack"
	"reflect"
	"testing"
)

func TestStack_IsEmptyWhenNothingHasBeenPushed(t *testing.T) {
	s := stack.NewStack[int]()

	if s.IsEmpty() != true {
		t.Errorf("should be empty")
	}
}

func TestStack_Peek(t *testing.T) {
	type want struct {
		data int
		ok   bool
	}
	type testCase struct {
		elems []int
		want  want
	}

	testTable := []testCase{
		{elems: []int{}, want: want{0, false}},
		{elems: []int{-1}, want: want{-1, true}},
		{elems: []int{1}, want: want{1, true}},
		{elems: []int{1, 2}, want: want{2, true}},
		{elems: []int{1, 2, 3}, want: want{3, true}},
	}

	for _, tc := range testTable {
		s := stack.NewStack[int]()
		for _, e := range tc.elems {
			s.Push(e)
		}

		got, ok := s.Peek()
		if ok != tc.want.ok {
			t.Errorf("expected ok to be %t but got %t", tc.want.ok, ok)
		}

		if got != tc.want.data {
			t.Errorf("expected data to be %d but got %d", tc.want.data, got)
		}
	}
}
func TestStack_PeekDoesntAlterTheStack(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	p1, _ := s.Peek()
	p2, _ := s.Peek()
	if p1 != p2 {
		t.Errorf("expected %d and %d to be equal", p1, p2)
	}
}

func TestStack_PopRemovesTheTopElement(t *testing.T) {
	type want struct {
		data int
		ok   bool
	}

	type testCase struct {
		name  string
		elems []int
		pops  int
		peek  want
	}

	testCases := []testCase{
		{
			name:  "pushing 0, popping 1",
			elems: []int{},
			pops:  1,
			peek:  want{0, false},
		},
		{
			name:  "pushing 3, popping 0",
			elems: []int{1, 2, 3},
			pops:  0,
			peek:  want{3, true},
		},
		{
			name:  "pushing 3, popping 1",
			elems: []int{1, 2, 3},
			pops:  1,
			peek:  want{2, true},
		},
		{
			name:  "pushing 3, popping 2",
			elems: []int{1, 2, 3},
			pops:  2,
			peek:  want{1, true},
		},
		{
			name:  "pushing 3, popping 3",
			elems: []int{1, 2, 3},
			pops:  3,
			peek:  want{0, false},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := stack.NewStack[int]()

			for _, e := range tc.elems {
				s.Push(e)
			}

			for i := 0; i < tc.pops; i++ {
				s.Pop()
			}

			top, ok := s.Peek()
			if ok != tc.peek.ok {
				t.Errorf("Peek() should have ok value of %t, got %t", tc.peek.ok, ok)
			}

			if top != tc.peek.data {
				t.Errorf("Peek() should have value of %d, got %d", tc.peek.data, top)
			}
		})
	}
}

func TestStack_ToSlice(t *testing.T) {
	s := stack.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	got := s.ToSlice()
	want := []int{3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
