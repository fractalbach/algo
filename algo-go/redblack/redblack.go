package redblack

import "fmt"

// Tree is a Red-Black Tree. Create one by using NewTree[T comparable](value T)
// to initialize the root of the tree.
type Tree[T comparable] struct {
	root *node[T]
}

// node in a red-black tree
type node[T comparable] struct {
	parent *node[T]
	left   *node[T]
	right  *node[T]
	color  color
	value  T
}

type color int

const (
	red color = iota
	black
)

// NewTree creates and initializes the root of a new Red-Black Tree. This will
// also determine the type of values that the tree contains. Since Red-Black
// trees are a type of binary search tree, the values must be comparable.
func NewTree[T comparable](value T) *Tree[T] {
	return &Tree[T]{
		root: &node[T]{nil, nil, nil, black, value},
	}
}

func (t *Tree[T]) leftRotate(x *node[T]) {
	y := x.right // tree re-configuration
	x.right = y.left
	y.left = x

	y.parent = x.parent // update links between parents and children
	if x.right != nil {
		x.right.parent = x
	}
	if y.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

func (t *Tree[T]) rightRotate(y *node[T]) {

}

func Hello() string {
	t := NewTree(5)
	fmt.Printf("%#v\n", t)
	return "hello"
}
