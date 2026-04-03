package main

// import (
// 	"fmt"
// )

func main() {
	cards := newDeck()
	// firstHand, remainingDeck := deal(cards, 5)
	// firstHand.print()
	// remainingDeck.print()
	// cardsString := cards.toString()
	// fmt.Println(cardsString)
	cards.saveToFile("my_cards")	
}


