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
	_, err := hello("")

	if err == nil {
		t.Fatal("an empty name should return an error")
	}

	expectedGreeting := "Hail, Doga! Well met!"
	actualGreeting, err := hello("Doga")

	if err != nil {
		t.Fatal(err)
	}

	if actualGreeting != expectedGreeting {
		t.Errorf("expected '%v'; got '%v'", expectedGreeting, actualGreeting)
	}

	expectedGreetings := []string{
		"Hi, Ji-an. Welcome!",
		"Hail, Hiroto! Well met!",
		"Hail, Somchai! Well met!",
	}

	actualGreetings := randomGreetings("Ji-an", "Hiroto", "Somchai")

	if !reflect.DeepEqual(actualGreetings, expectedGreetings) {
		t.Errorf("expected '%v'; got '%v'", expectedGreetings, actualGreetings)
	}
}
