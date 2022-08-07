package redblack

import (
	"testing"
)

func TestHello(t *testing.T) {
	tree := NewTree(5)
	for i := 0; i < 10; i++ {
		tree.Insert(i)
	}
	t.Logf("tree: %v\n", tree)
	t.Logf("%v", tree.string())
}

func TestRotation(t *testing.T) {
	tree := NewTree(5)
	for _, v := range []int{2, 1, 3, 7} {
		tree.Insert(v)
	}
	t.Logf("tree starts as: %v", tree.string())
	tree.rightRotate(tree.root)
	t.Logf("do rightRotate: %v", tree.string())
	tree.leftRotate(tree.root)
	t.Logf("do leftRotate : %v", tree.string())
}

// (((1) 2 (3)) 5 (7))
