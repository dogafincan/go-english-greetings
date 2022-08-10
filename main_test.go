package main

import "testing"

func TestHello(t *testing.T) {
	greeting, error := hello("Doga")
	if greeting != "Hello, Doga!" {
		t.Fatalf("%v", error)
	}
}
