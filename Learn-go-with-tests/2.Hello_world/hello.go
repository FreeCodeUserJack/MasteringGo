package main

import "fmt"

const englishHelloPrefix = "hello "
const englishHelloSuffix = "world"

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + englishHelloSuffix
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Jack"))
}