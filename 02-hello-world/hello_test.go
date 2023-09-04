package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Bob Bobson")
	expected := "Hello, Bob Bobson!"

	if result != expected {
		t.Errorf("got %q, but want %q", result, expected)
	}
}
