package main

import (
	"deck/deck-of-cards"
)

func main() {
	newDeck := deck_of_cards.CreateDeck52()
	deck_of_cards.PrintDeck(newDeck)
	newDeck = deck_of_cards.AddJokers(newDeck, 3)
	deck_of_cards.ShuffleDeck(newDeck)
	deck_of_cards.PrintDeck(newDeck)
	deck_of_cards.SortDeck(newDeck)
	deck_of_cards.PrintDeck(newDeck)
	newDeck = deck_of_cards.DeleteCard(newDeck, deck_of_cards.Joker)
	deck_of_cards.PrintDeck(newDeck)

	multiplyDeck := deck_of_cards.CreateMultiplyDeck(2)
	deck_of_cards.PrintDeck(multiplyDeck)
	multiplyDeck = deck_of_cards.DeleteCard(multiplyDeck, deck_of_cards.Ace)
	deck_of_cards.PrintDeck(multiplyDeck)

}
