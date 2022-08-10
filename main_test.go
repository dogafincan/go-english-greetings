package main

import "testing"

func TestHello(t *testing.T) {
	greeting, error := hello("Doga")
	if greeting != "Hello, Doga!" {
		t.Fatalf("%v", error)
	}

	_, error = hello("")
	if error == nil {
		t.Fatalf("return an error when name is an empty string")
	}
}
