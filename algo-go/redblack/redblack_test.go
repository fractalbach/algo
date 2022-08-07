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
