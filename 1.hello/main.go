package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "
const defaultReceiver = "World"

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultReceiver
	}

	greeting := englishGreeting
	if lang == spanish {
		greeting = spanishGreeting
	}
	if lang == french {
		greeting = frenchGreeting
	}

	return greeting + name
}
