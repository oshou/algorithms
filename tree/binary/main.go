package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{
		value: v,
		left:  nil,
		right: nil,
	}
}

func (n *Node) Insert(v int) {
	if n == nil {
		n = NewNode(v)
		return
	}

	switch {
	case v < n.value:
		if n.left == nil {
			n.left = NewNode(v)
			return
		}
		n.left.Insert(v)
	case v > n.value:
		if n.right == nil {
			n.right = NewNode(v)
			return
		}
		n.right.Insert(v)
	}
}

func (n *Node) Delete(v int) *Node {
	if n == nil {
		return nil
	}

	switch {
	case v < n.value:
		n.left = n.left.Delete(v)
	case v > n.value:
		n.right = n.right.Delete(v)
	default:
		switch {
		case n.left == nil:
			return n.right
		case n.right == nil:
			return n.left
		default:
			tmp := n.right.Min()
			n.value = tmp.value
			n.right = n.right.Delete(tmp.value)
		}
	}
	return n
}

func (n *Node) Min() *Node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.Min()
	} else {
		return n
	}
}

func (n *Node) Inorder() {
	if n != nil {
		n.left.Inorder()
		fmt.Println(n.value)
		n.right.Inorder()
	}

}

func main() {
	root := &Node{1, nil, nil}
	root.Insert(3)
	root.Insert(5)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Inorder()
	root.Delete(5)
	root.Inorder()
}
