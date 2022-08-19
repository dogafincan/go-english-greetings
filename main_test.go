package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Doga!"
	actual, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected '%v'. Got '%v'", expected, actual)
	}
}
