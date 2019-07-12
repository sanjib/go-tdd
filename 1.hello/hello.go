package hello

const spanish = "Spanish"
const french = "French"
const bengali = "Bengali"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "
const bengaliGreeting = "Slamalaikum, "

const defaultReceiver = "World"

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultReceiver
	}
	return greeting(lang) + name
}

func greeting(lang string) (greeting string) {
	switch lang {
	case spanish:
		greeting = spanishGreeting
	case french:
		greeting = frenchGreeting
	case bengali:
		greeting = bengaliGreeting
	default:
		greeting = englishGreeting
	}
	return
}
