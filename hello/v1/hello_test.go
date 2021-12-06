package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello,Shane Lea"

	if got != want {
		t.Errorf("got : '%q', want: '%q'", got, want)
	}
}
