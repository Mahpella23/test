package main

import "testing"

func TestHello(t *testing.T) {
	s := Hello()

	if s != "Hello" {
		t.Error("Function Hello returns wrong value, expected 'Hello', acquired " + s)
	}
}
