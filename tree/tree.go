package tree

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	data  T
	left  *node[T]
	right *node[T]
}

func NewNode[T constraints.Ordered](data T) *node[T] {
	return &node[T]{data, nil, nil}
}

func (root *node[T]) Insert(data T) *node[T] {
	if root == nil {
		return NewNode[T](data)
	}

	if data <= root.data {
		root.left = root.left.Insert(data)
	} else {
		root.right = root.right.Insert(data)
	}

	return root
}

func (root *node[T]) Contains(data T) bool {
	switch {
	case root == nil:
		return false
	case data == root.data:
		return true
	case data < root.data:
		return root.left.Contains(data)
	case data > root.data:
		return root.right.Contains(data)
	default:
		return false
	}
}

func (root *node[T]) ToSlice() []T {
	if root == nil {
		return []T{}
	}

	result := []T{}
	result = append(result, root.left.ToSlice()...)
	result = append(result, root.data)
	result = append(result, root.right.ToSlice()...)
	return result
}
