package main

import "fmt"

var langPrefixes = map[string]string{
	"en": "Hello",
	"es": "Hola",
}

func main() {
	fmt.Println(Hello("en", "Joe"))
}

func Hello(lang, name string) string {
	prefix, exists := langPrefixes[lang]

	if exists {
		return buildGreeting(prefix, name)
	} 

	return buildGreeting(langPrefixes["en"], name)
}

func buildGreeting(prefix, name string) string {
	if len(name) == 0 {
		return prefix + ", World"
	}

	return prefix + ", " + name
}