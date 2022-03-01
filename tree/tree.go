package tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data, nil, nil}
}

func (root *Node) Insert(data int) *Node {
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

func (root *Node) Contains(data int) bool {
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

func (root *Node) ToSlice() []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	result = append(result, root.left.ToSlice()...)
	result = append(result, root.data)
	result = append(result, root.right.ToSlice()...)
	return result
}
