package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("when valid name", func(t *testing.T) {
			got := Hello("Joe")
			want := "Hello, Joe"

			assertEqual(t, got, want)
	})
	t.Run("when empty string", func(t *testing.T) {
			got := Hello("")
			want := "Hello, World"

			assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
