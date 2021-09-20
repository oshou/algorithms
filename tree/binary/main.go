package main

import "fmt"

type Node interface {
	Insert(v int)
	Delete(v int)
	Search(v int) *Node
}

type node struct {
	value int
	left  *node
	right *node
}

func NewNode(v int) *node {
	return &node{
		value: v,
		left:  nil,
		right: nil,
	}
}

func (n *node) Insert(v int) {
	if n == nil {
		n = NewNode(v)
	}

	switch {
	case n.value > v:
		if n.left == nil {
			n.left = NewNode(v)
		} else {
			n.left.Insert(v)
		}
	case n.value < v:
		if n.right == nil {
			n.right = NewNode(v)
		} else {
			n.right.Insert(v)
		}
	}
}

func (n *node) Delete(v int) {
	if n == nil {
		return
	}

	switch {
	case n.value > v:
		n.left.Delete(v)
	case n.value < v:
		n.right.Delete(v)
	default:
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.left == nil {
			*n = *n.right
		} else if n.right == nil {
			*n = *n.left
		} else {
			min := n.right.SearchMin()
			n.value = min.value
			n.right.Delete(min.value)
		}
	}
}

func (n *node) Search(v int) *node {
	if n == nil {
		return nil
	}
	switch {
	case n.value == v:
		return n
	case n.value > v:
		return n.left.Search(v)
	default:
		return n.right.Search(v)
	}
}

func (n *node) SearchMin() *node {
	if n == nil {
		return nil
	}
	if n.left != nil {
		return n.left.SearchMin()
	} else {
		return n
	}
}

func (n *node) inorder() {
	if n != nil {
		n.left.inorder()
		fmt.Println(n.value)
		n.right.inorder()
	}
}

func (n *node) preorder() {
	if n != nil {
		fmt.Println(n.value)
		n.left.inorder()
		n.right.inorder()
	}
}

func (n *node) postorder() {
	if n != nil {
		n.left.inorder()
		n.right.inorder()
		fmt.Println(n.value)
	}
}

func main() {
	root := NewNode(5)
	root.Insert(1)
	root.Insert(9)
	root.Insert(3)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)

	root.Delete(5)
	fmt.Println("r:", root.value)
	fmt.Println("r->l", root.left.value)
	fmt.Println("r->l->r", root.left.right.value)
	fmt.Println("r->l->r->l", root.left.right.left.value)
	fmt.Println("r->l->r->r", root.left.right.right.value)
	fmt.Println("r->r", root.right.value)
	fmt.Println("r->r->l", root.right.left.value)

	fmt.Println("inorder")
	root.inorder()

	fmt.Println("preorder")
	root.preorder()

	fmt.Println("postorder")
	root.postorder()
}
