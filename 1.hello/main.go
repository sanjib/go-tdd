package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const bengali = "Bengali"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "
const bengaliGreeting = "Slamalaikum, "

const defaultReceiver = "World"

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultReceiver
	}

	greeting := englishGreeting

	switch lang {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	case bengali:
		greeting = bengaliGreeting
	}

	return greeting + name
}
