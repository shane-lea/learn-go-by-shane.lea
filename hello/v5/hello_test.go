package main

import "testing"

func TestHello(t *testing.T) {

	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: '%q', want: '%q'", got, want)
		}

	}

	t.Run("test-001", func(t *testing.T) {
		got := Hello("Shane Lea")
		want := "Hello, Shane Lea"
		assert(t, got, want)
	})

	t.Run("test-002", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Shane Lea"
		assert(t, got, want)
	})

}
