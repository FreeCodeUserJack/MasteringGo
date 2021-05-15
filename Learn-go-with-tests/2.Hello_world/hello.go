package main

import "fmt"

const englishHelloPrefix = "hello "
const spanishHelloPrefix = "hola "
const frenchHelloPrefix = "bonjour "
const southernHelloPrefix = "howdy "

const spanish = "Spanish"
const french = "French"
const southern = "Southern"

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// } else if language == french {
	// 	return frenchHelloPrefix + name
	// }
	// return englishHelloPrefix + name

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case southern:
		prefix = southernHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Jack", ""))
}