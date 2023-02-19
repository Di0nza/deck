package main

import (
	deck "github.com/Di0nza/deckOfCards/pkg/deck"
)

func main() {
	newDeck := deck.NewDeck()
	deck.PrintDeck(newDeck)
	newDeck = deck.AddJokers(newDeck, 2)
	deck.ShuffleDeck(newDeck)
	deck.PrintDeck(newDeck)
	deck.SortDeck(newDeck)
	deck.PrintDeck(newDeck)
	newDeck = deck.DeleteCard(newDeck, deck.Ace)
	deck.PrintDeck(newDeck)

	multiplyDeck := deck.CreateMultiplyDeck(2)
	deck.PrintDeck(multiplyDeck)
	multiplyDeck = deck.DeleteCard(multiplyDeck, deck.Ace)
	deck.PrintDeck(multiplyDeck)
}
