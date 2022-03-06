package tree

type node struct {
	data  int
	left  *node
	right *node
}

func NewNode(data int) *node {
	return &node{data, nil, nil}
}

func (root *node) Insert(data int) *node {
	if root == nil {
		return NewNode(data)
	}

	if data <= root.data {
		root.left = root.left.Insert(data)
	} else {
		root.right = root.right.Insert(data)
	}

	return root
}

func (root *node) Contains(data int) bool {
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

func (root *node) ToSlice() []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.left.ToSlice()...)
	result = append(result, root.data)
	result = append(result, root.right.ToSlice()...)
	return result
}
