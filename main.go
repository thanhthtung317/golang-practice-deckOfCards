package main

import "fmt"

func main() {
	cards := newDeck()

	// // cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	fmt.Println([]byte(cards.toString()))
}

func newCard() string {
	return "five of diamonds"
}