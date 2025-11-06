package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return  spanishHelloPrefix + name
	}
	if language == french {
		return  frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	name := ""
	language := ""
	fmt.Println(Hello(name, language));
}