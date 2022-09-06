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

func RandomGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("expected a non-empty string")
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
		return []string{}, errors.New("expected at least one name")
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
