package hello

import "testing"

func TestHello(t *testing.T) {
	s := Hello()
	if s != "Hello, world.1" {
		t.Errorf("error")
	}
}