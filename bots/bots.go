package bots

import "fmt"

type bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func PrintGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (EnglishBot) getGreeting() string {
	return "Hi there!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}
