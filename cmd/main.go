package main

import (
	"deckOfCards/pkg/deck"
)

func main() {
	newDeck := deck.CreateDeck52()
	deck.PrintDeck(newDeck)
	newDeck = deck.AddJokers(newDeck, 3)
	deck.ShuffleDeck(newDeck)
	deck.PrintDeck(newDeck)
	deck.SortDeck(newDeck)
	deck.PrintDeck(newDeck)
	newDeck = deck.DeleteCard(newDeck, deck.Joker)
	deck.PrintDeck(newDeck)

	multiplyDeck := deck.CreateMultiplyDeck(2)
	deck.PrintDeck(multiplyDeck)
	multiplyDeck = deck.DeleteCard(multiplyDeck, deck.Ace)
	deck.PrintDeck(multiplyDeck)
}
