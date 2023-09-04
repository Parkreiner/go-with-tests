package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Says hello to people", func(t *testing.T) {
		result := Hello("Bob Bobson")
		expected := "Hello, Bob Bobson!"

		if result != expected {
			t.Errorf("got %q, but want %q", result, expected)
		}
	})

	t.Run("Says 'Hello, world!' when an empty string is provided", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world!"

		if result != expected {
			t.Errorf("got %q, but want %q", result, expected)
		}
	})
}
