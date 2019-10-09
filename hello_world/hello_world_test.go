package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world")
	want := "hello, world"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestHello2(t *testing.T) {
	got := Hello("george")
	want := "hello, george"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
