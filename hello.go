package main

import "fmt"

func main() {
	fmt.Println(Hello("Joe"))
}

func Hello(name string) string {
	return "Hello, " + name
}