package main

import "testing"

func assertCorrectMessage(t testing.TB, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("got %q, but want %q", received, expected)
	}
}

func TestHello(t *testing.T) {
	t.Run("Says hello to people", func(t *testing.T) {
		result := Hello("Bob Bobson")
		expected := "Hello, Bob Bobson!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("Says 'Hello, world!' when an empty string is provided", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, world!"
		assertCorrectMessage(t, result, expected)
	})
}
