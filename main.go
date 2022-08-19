package main

import (
	"errors"
	"fmt"
	"log"
)

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name cannot be an empty string")
	}

	return fmt.Sprintf("Hello, %v!", name), nil
}

func main() {
	greeting, err := hello("Doga")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}
