package main

import "fmt"

func main() {
	cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//fmt.Println()
	//remainingCards.print()
	cards.saveToFile("my_cards")
	newDeckFromFile("my_cards").print()
	fmt.Printf("\n\n\n")
	cards.shuffle()
}
