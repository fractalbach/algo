package redblack

import (
	"testing"
)

func TestHello(t *testing.T) {
	tree := NewTree(5)
	t.Logf("tree: %v\n", tree)
}
