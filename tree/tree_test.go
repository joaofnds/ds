package tree_test

import (
	"ds/tree"
	"reflect"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	testTable := []struct {
		inserts []int
		want    []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
	}

	for _, tc := range testTable {
		tree := tree.NewNode(tc.inserts[0])

		for _, e := range tc.inserts[1:] {
			tree.Insert(e)
		}

		got := tree.ToSlice()

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestNode_Contains(t *testing.T) {
	testTable := []struct {
		inserts []int
		n       int
		want    bool
	}{
		{[]int{1}, 1, true},
		{[]int{1}, 0, false},

		{[]int{1, 2}, 2, true},
		{[]int{2, 1}, 1, true},

		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 3, true},

		{[]int{2, 1, 3}, 0, false},
		{[]int{2, 1, 3}, 4, false},
	}

	for _, tc := range testTable {
		tree := tree.NewNode(tc.inserts[0])

		for _, e := range tc.inserts[1:] {
			tree.Insert(e)
		}

		got := tree.Contains(tc.n)

		if got != tc.want {
			t.Errorf("after inserting %v, expected Contains(%d) to be %t, but got %t", tc.inserts, tc.n, tc.want, got)
		}
	}
}
