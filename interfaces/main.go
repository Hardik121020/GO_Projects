package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishbot struct{}
type spanishbot struct{}

func (eb englishbot) getGreeting() string {
	return "Hello There"
}

func (sb spanishbot) getGreeting() string {
	return "Hola !"
}

func PrintGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	eb := englishbot{}
	sb := spanishbot{}
	PrintGreeting(eb)
	PrintGreeting(sb)
}
