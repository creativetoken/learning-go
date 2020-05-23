package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}
type hindiBot struct{}

func main() {

	rb := englishBot{}
	sb := spanishBot{}
	hb := hindiBot{}
	printGreeting(rb)
	printGreeting(sb)
	printGreeting(hb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {

	return "Hi There"
}

func (spanishBot) getGreeting() string {

	return "Hola!"
}

func (hindiBot) getGreeting() string {

	return "Kaise ho"
}
