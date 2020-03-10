package hello

import "testing"

func TestHello(t *testing.T) {
	s := Hello()
	if s != "Hello, world." {
		t.Errorf("error")
	}
}