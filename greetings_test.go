package greetings

import (
	"errors"
	"math/rand"
	"reflect"
	"testing"
)

func init() {
	rand.Seed(1)
}

func TestRandomGreeting(t *testing.T) {
	expected := "Hail, Doga! Well met!"
	actual, err := RandomGreeting("Doga")

	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}

func TestRandomGreetingEmptyString(t *testing.T) {
	_, err := RandomGreeting("")

	if !errors.Is(err, ErrEmptyNameArgument) {
		t.Error("an empty string should return an ErrEmptyNameArgument")
	}
}

func TestRandomGreetings(t *testing.T) {
	expected := []string{
		"Hi, Ji-an. Welcome!",
		"Hail, Hiroto! Well met!",
		"Hail, Somchai! Well met!",
	}

	actual, err := RandomGreetings("Ji-an", "Hiroto", "Somchai")

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v'; got '%v'", expected, actual)
	}
}

func TestRandomGreetingsEmptyArray(t *testing.T) {
	_, err := RandomGreetings()

	if !errors.Is(err, ErrEmptyNamesArgument) {
		t.Error("no names entered should return an ErrEmptyNamesArgument")
	}
}

func TestRandomGreetingsEmptyString(t *testing.T) {
	_, err := RandomGreetings("")

	if !errors.Is(err, ErrEmptyNameArgument) {
		t.Error("one of the names being an empty string should return an ErrEmptyNameArgument")
	}
}
