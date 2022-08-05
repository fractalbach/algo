package redblack

import "golang.org/x/exp/constraints"

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
	}
	z.left = nil
	z.right = nil
	z.color = red
	t.insertFixup()
}

func (t *Tree[T]) insertFixup() {

}
