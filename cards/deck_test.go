package main

import (
	"os"
	"testing"
)

func TestDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected 52 but got %d", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Wrong First Card")
	}

	if cards[len(cards)-1] != "king of Clubs" {
		t.Errorf("Wrong Last Card")
	}
}

func Test_fileCreation(t *testing.T) {
	d := newDeck()

	os.Remove("my_cards")
	d.saveToFile("my_cards")
	d2 := newDeckFromFile("my_cards")
	if len(d2) != 52 {
		t.Errorf("There is some error")
	}
	os.Remove("my_cards")
}
