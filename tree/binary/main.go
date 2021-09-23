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
		return
	}

	switch {
	case n.value > v:
		if n.left == nil {
			n.left = NewNode(v)
			return
		}
		n.left.Insert(v)
	case n.value < v:
		if n.right == nil {
			n.right = NewNode(v)
			return
		}
		n.right.Insert(v)
	}
}

func (n *node) Delete(v int) {
	if n == nil {
		return
	}

	switch {
	case n.value > v:
		if n.left.value == v {
			n.left = nil
		}
		n.left.Delete(v)
	case n.value < v:
		if n.right.value == v {
			n.left = nil
		}
		n.right.Delete(v)
	default:
		if n.left == nil {
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
		fmt.Println(nil)
		return nil
	}
	switch {
	case n.value == v:
		return n
	case n.value > v:
		if n.left == nil {
			return nil
		}
		return n.left.Search(v)
	default:
		if n.right == nil {
			return nil
		}
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

func (n *node) InOrder() {
	if n != nil {
		n.left.InOrder()
		fmt.Println(n.value)
		n.right.InOrder()
	}
}

func (n *node) PreOrder() {
	if n != nil {
		fmt.Println(n.value)
		n.left.PreOrder()
		n.right.PreOrder()
	}
}

func (n *node) PostOrder() {
	if n != nil {
		n.left.PostOrder()
		n.right.PostOrder()
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

	fmt.Printf("t: %v\n", root)
	fmt.Printf("t->l: %v\n", root.left)
	fmt.Printf("t->l->r: %v\n", root.left.right)
	fmt.Printf("t->l->r->l: %v\n", root.left.right.left)
	fmt.Printf("t->l->r->r: %v\n", root.left.right.right)
	fmt.Printf("t->r: %v\n", root.right)

	fmt.Println("search:", root.Search(5))
	fmt.Println("search:", root.Search(6))

	fmt.Println("--- InOrder ---")
	root.InOrder()

	fmt.Println("--- PreOrder ---")
	root.PreOrder()

	fmt.Println("--- PostOrder ---")
	root.PostOrder()
}
