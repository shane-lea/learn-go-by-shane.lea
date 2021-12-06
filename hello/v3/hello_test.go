package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("这是一个分组测试的名字:say hello to myself", func(t *testing.T) {
		got := Hello("Shane Lea")
		want := "Hello, Shane Lea"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Shane Lea"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}

	})
}
