package main

import "testing"

func TestHello(t *testing.T) {
	_, err := hello("")

	if err == nil {
		t.Fatal("An empty name should return an error")
	}

	expected := "Hello, Doga!"
	actual, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("Expected '%v'. Got '%v'", expected, actual)
	}
}
