package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected deck length of 20, but got %v", len(d))
	}
	if d[0] != "spades of ace"{
		t.Errorf("expected first card of ace of spades but got %v", d[0])
	}
	if d[len(d) - 1] != "clubs of four"{
		t.Errorf("expected first card of four of clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16{
		t.Errorf("expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}