package greetings

import (
	"math/rand"
	"reflect"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestRandomGreeting(t *testing.T) {
	expected := "Hail, Doga! Well met!"
	actual, err := randomGreeting("Doga")

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}

func TestRandomGreetingEmpty(t *testing.T) {
	greeting, err := randomGreeting("")

	if greeting != "" || err == nil {
		t.Error("an empty name should return an error")
	}
}

func TestRandomGreetings(t *testing.T) {
	expected := []string{
		"Hi, Ji-an. Welcome!",
		"Hail, Hiroto! Well met!",
		"Hail, Somchai! Well met!",
	}

	actual, err := randomGreetings("Ji-an", "Hiroto", "Somchai")

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}

func TestRandomGreetingsEmpty(t *testing.T) {
	greetings, err := randomGreetings()

	if !reflect.DeepEqual(greetings, []string{}) || err == nil {
		t.Error("an empty name should return an error")
	}
}
