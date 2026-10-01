package main

type deck []string

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
