package main

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

func randomGreeting() string {
	greetings := [3]string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return greetings[rand.Intn(len(greetings))]
}

func randomGreetings(names ...string) ([]string, error) {
	if len(names) == 0 {
		return []string{}, errors.New("no names found")
	}

	var greetings []string

	for _, name := range names {
		greeting, err := hello(name)

		if err != nil {
			log.Fatal(err)
		}

		greetings = append(greetings, greeting)
	}

	return greetings, nil
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be an empty string")
	}

	return fmt.Sprintf(randomGreeting(), name), nil
}

func main() {
	greeting, err := hello("Doga")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)

	greetings, err := randomGreetings("Ji-an", "Hiroto", "Somchai")

	if err != nil {
		log.Fatal(err)
	}

	for _, greeting := range greetings {
		fmt.Println(greeting)
	}
}
