package main

import "testing"

func TestHello(t *testing.T) {
	s := Hello()

	if s != "Hello, World" {
		t.Error("Function HelloWorld returns wrong value, expected 'Hello', acquired " + s)
	}
}
