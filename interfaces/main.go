package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}


func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

// Cant have two functions with the same name and different parameters in Go. The above code will not compile because of the duplicate function name `printGreeting`. You can use interfaces to achieve polymorphism instead.
/* func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sp spanishBot) {
	fmt.Println(sp.getGreeting())
} */

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

