package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"spades", "diamond", "heart", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}
 
func (d deck) print() deck {
	for i, card := range d {
		fmt.Println(i, card)
	}
	return []string{"kk"}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string{
	return strings.Join([]string(d),",")
}