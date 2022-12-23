package main

import "fmt"

// import (
// 	"fmt"
// )

func main() {
	cards := newDeck()
	cards.shuffle()

	// // cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println([]byte(cards.toString()))
	// cards.saveToFile("my_cards")
	fmt.Println(cards.print())
}

func newCard() string {
	return "five of diamonds"
}

