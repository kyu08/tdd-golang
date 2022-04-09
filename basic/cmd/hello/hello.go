package hello

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const defaultName = "World"

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}
	return helloPrefix(lang) + name
}

func helloPrefix(lang string) string {
	switch lang {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
	}
}
