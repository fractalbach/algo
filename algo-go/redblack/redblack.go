package redblack

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

// Tree is a Red-Black Tree. Create one by using NewTree[T comparable](value T)
// to initialize the root of the tree.
type Tree[T constraints.Ordered] struct {
	root *node[T]
}

// node in a red-black tree
type node[T constraints.Ordered] struct {
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
func NewTree[T constraints.Ordered](value T) *Tree[T] {
	return &Tree[T]{
		root: &node[T]{nil, nil, nil, black, value},
	}
}

// assumes x.right is not nil
func (t *Tree[T]) leftRotate(x *node[T]) {
	y := x.right // declares node y

	x.right = y.left // y's left subtree becomes x's right subtree
	y.left = x
	if x.right != nil {
		x.right.parent = x
	}

	y.parent = x.parent // link x's parent to y
	if x.parent == nil {
		t.root = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x // x becomes y's left subtree
	x.parent = y
}

// assumes y.left is not nil
func (t *Tree[T]) rightRotate(y *node[T]) {
	x := y.left // declares node x

	y.left = x.right // x's right subtree becomes y's left subtree
	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent // link y's parent to x
	if y.parent == nil {
		t.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	x.right = y // y becomes x's right subtree
	y.parent = x
}

// Insert an element into the red-black tree in O(lg n) time complexity. The
// tree will automatically balance itself to ensure this.
func (t *Tree[T]) Insert(value T) {
	var x *node[T] = t.root                 // 'current' node as we traverse
	var y *node[T] = nil                    // parent of node to insert
	var z *node[T] = &node[T]{value: value} // node to insert
	for x != nil {
		y = x
		if z.value < y.value {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		t.root = z
	} else if z.value < y.value {
		y.left = z
	} else {
		y.right = z
	}
	z.left = nil
	z.right = nil
	z.color = red
	t.insertFixup(z)
}

func (t *Tree[T]) insertFixup(z *node[T]) {
	for z.parent != nil && z.parent.color == red {

		// Find the node z's uncle, and determine what side it's on. Start by
		// assuming it's uncle is on left, then fix if needed.
		var uncle *node[T] = z.parent.parent.left
		var isUncleOnRight bool = false
		if z.parent == uncle {
			uncle = z.parent.parent.right
			isUncleOnRight = true
		}

		// Case 1: When uncle is red, all we need to do is change the colorings
		// of some nodes, and move the z pointer.
		if (uncle != nil) && (uncle.color == red) {
			z.parent.color = black
			uncle.color = black
			z = z.parent.parent
			z.color = red
			continue
		}

		// Case 2/3: When uncle is black (or nil), then we need to do rotations,
		// the side that the uncle is on is important.
		if isUncleOnRight {
			if z == z.parent.right { // Case 2
				z = z.parent
				t.leftRotate(z)
			} else { // Case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.rightRotate(z.parent.parent)
			}

		} else { // Uncle on Left
			if z == z.parent.left { // Case 2
				z = z.parent
				t.rightRotate(z)
			} else { // Case 3
				z.parent.color = black
				z.parent.parent.color = red
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = black
}

// returns a string containing a nested-parentheses representation of the tree.
// The nodes are printed using pre-order traversal, so the general format for
// each node will be (value(left)(right)). A more complete example looks like
// (5(2(1)(3))(7)).
func (t *Tree[T]) string() string {
	var b strings.Builder
	buildString(&b, t.root)
	return b.String()
}

func buildString[T constraints.Ordered](b *strings.Builder, n *node[T]) {
	if n == nil {
		return
	}
	b.WriteRune('(')
	fmt.Fprintf(b, "%v", n.value)
	buildString(b, n.left)
	buildString(b, n.right)
	b.WriteRune(')')
}
