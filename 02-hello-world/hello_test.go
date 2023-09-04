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
		result := Hello("Bob Bobson", "en")
		expected := "Hello, Bob Bobson!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("Says 'Hello, world!' when an empty string is provided", func(t *testing.T) {
		result := Hello("", "en")
		expected := "Hello, world!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("Greets the user in Spanish if locale is Spanish", func(t *testing.T) {
		result := Hello("Don Quixote", "es")
		expected := "Hola, Don Quixote!"
		assertCorrectMessage(t, result, expected)
	})

	t.Run("Defaults to English if unknown locale is used", func(t *testing.T) {
		result := Hello("Blobfish", "zh")
		expected := "Hello, Blobfish!"
		assertCorrectMessage(t, result, expected)
	})
}
