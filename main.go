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

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be an empty string")
	}

	return fmt.Sprintf(randomGreeting(), name), nil
}

func main() {
	log.SetPrefix("hello: ")
	log.SetFlags(0)

	greeting, err := hello("Doga")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}
