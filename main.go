package main

import (
	"fmt"
	"os"
)

func main() {
	cards := newDeck()
	// firstHand, remainingDeck := deal(cards, 5)
	// firstHand.print()
	// remainingDeck.print()
	cardsString := cards.toString()
	fmt.Println(cardsString)
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
