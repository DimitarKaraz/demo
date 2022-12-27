package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Joe"))
}

func Hello(name string) string {
	if len(name) == 1 {
		return englishHelloPrefix + "World"
	}

	return englishHelloPrefix + name
}