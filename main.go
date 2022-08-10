package main

import "fmt"

func hello(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be an empty string")
	}

	return "Hello, " + name + "!", nil
}

func main() {
	fmt.Println(hello("Doga"))
}
