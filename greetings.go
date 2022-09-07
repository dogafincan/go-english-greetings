package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ErrEmptyNameArgument  = errors.New("expected a non-empty string")
	ErrEmptyNamesArgument = errors.New("expected at least one name")
)

func RandomGreeting(name string) (string, error) {
	if name == "" {
		return "", ErrEmptyNameArgument
	}

	greetings := [3]string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	randomGreeting := greetings[rand.Intn(len(greetings))]

	return fmt.Sprintf(randomGreeting, name), nil
}

func RandomGreetings(names ...string) ([]string, error) {
	if len(names) == 0 {
		return []string{}, ErrEmptyNamesArgument
	}

	var greetings []string

	for _, name := range names {
		greeting, err := RandomGreeting(name)

		if err != nil {
			return []string{}, err
		}

		greetings = append(greetings, greeting)
	}

	return greetings, nil
}
