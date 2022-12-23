package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) saveToFile(filename string, ) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)

		//Exit(0) means the program run successfully, beside that the program fail
		os.Exit(1)
	}

	return deck(strings.Split(string(bs),",")) 
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d{
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}