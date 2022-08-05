package redblack

import "testing"

func TestHello(t *testing.T) {
	want := "hello"
	got := Hello()
	if got != want {
		t.Errorf("Hello() = %q, wanted %q", got, want)
	}
}
