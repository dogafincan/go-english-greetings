package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestHello(t *testing.T) {
	expected := "Hail, Doga! Well met!"
	actual, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}

func TestHelloEmpty(t *testing.T) {
	greeting, err := hello("")

	if greeting != "" || err == nil {
		t.Fatal("an empty name should return an error")
	}
}

func TestRandomGreetings(t *testing.T) {
	expected := []string{
		"Hi, Ji-an. Welcome!",
		"Hail, Hiroto! Well met!",
		"Hail, Somchai! Well met!",
	}

	actual := randomGreetings("Ji-an", "Hiroto", "Somchai")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}
