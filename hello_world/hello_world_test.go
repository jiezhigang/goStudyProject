package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("say hello, george", func(t *testing.T) {
		got := Hello("george", "")
		want := "hello, george"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hell world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Chinese", func(t *testing.T) {
		got := Hello("中国", "Chinese")
		want := "你好，中国"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("so", "French")
		want := "Bonjour, so"

		assertCorrectMessage(t, got, want)
	})
}
