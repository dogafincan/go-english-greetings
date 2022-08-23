package main

import (
	"math/rand"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestHello(t *testing.T) {
	_, err := hello("")

	if err == nil {
		t.Fatal("an empty name should return an error")
	}

	expected := "Hail, Doga! Well met!"
	actual, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}
