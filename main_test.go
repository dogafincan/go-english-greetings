package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Doga!"
	actual, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actual != "Hello, Doga!" {
		t.Errorf("Expected '%v'. Got '%v'", expected, actual)
	}
}
