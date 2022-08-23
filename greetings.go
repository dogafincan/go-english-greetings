package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomGreeting(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be an empty string")
	}

	greetings := [3]string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	randomGreeting := greetings[rand.Intn(len(greetings))]

	return fmt.Sprintf(randomGreeting, name), nil
}

func randomGreetings(names ...string) ([]string, error) {
	if len(names) == 0 {
		return []string{}, errors.New("no names found")
	}

	var greetings []string

	for _, name := range names {
		greeting, err := randomGreeting(name)

		if err != nil {
			log.Fatal(err)
		}

		greetings = append(greetings, greeting)
	}

	return greetings, nil
}
