package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Lance")
	want := "Hello, Lance!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
