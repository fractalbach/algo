package redblack

import "fmt"

// tree is the root of a red-black tree, the root should always be colored black
type tree[T comparable] struct {
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

func newTree[T comparable](value T) *tree[T] {
	return &tree[T]{
		root: &node[T]{nil, nil, nil, black, value},
	}
}

func Hello() string {
	t := newTree(5)
	fmt.Printf("%#v\n", t)
	return "hello"
}
